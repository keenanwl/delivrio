package utils

import (
	"bytes"
	"context"
	"crypto/rand"
	"database/sql/driver"
	"delivrio.io/go/appconfig"
	"delivrio.io/go/carrierapis/labelutils"
	"delivrio.io/go/ent/user"
	"delivrio.io/go/viewer"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"gocloud.dev/gcerrors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/url"
	"os"
	"os/exec"
	"path"
	"regexp"
	"simonwaldherr.de/go/zplgfa"
	"strings"
	"time"

	"delivrio.io/go/ent"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"gocloud.dev/blob/fileblob"
	"golang.org/x/crypto/bcrypt"
)

var conf *appconfig.DelivrioConfig
var confSet = false

func Init(c *appconfig.DelivrioConfig) {
	if confSet {
		panic("utils: may not set config twice")
	}
	conf = c
	confSet = true
}

func ReadBody(body io.ReadCloser, toStruct interface{}) error {

	out, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(out, &toStruct)
	return err

}

// JoinPDFs joins 1 or more base64 encoded PDFs into a single PDF
func JoinPDFs(base64PDFs ...string) (string, error) {

	if len(base64PDFs) == 0 {
		return "", fmt.Errorf("0 PDFs may be joined")
	}

	mergers := make([]io.ReadSeeker, 0)
	for _, p := range base64PDFs {
		decoded, err := base64.StdEncoding.DecodeString(p)
		if err != nil {
			return "", err
		}
		mergers = append(mergers, bytes.NewReader(decoded))
	}

	res := bytes.NewBuffer(nil)

	err := api.MergeRaw(mergers, res, nil)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(res.Bytes()), nil

}

func Base64PDFToZPL(base64PDF string) (string, error) {

	img, err := Base64PDFToPNG(base64PDF)
	if err != nil {
		return "", err
	}

	// flatten image
	flat := zplgfa.FlattenImage(*img)

	// convert image to zpl compatible type
	gfimg := zplgfa.ConvertToZPL(flat, zplgfa.CompressedASCII)

	return gfimg, nil

}

func Base64PDFToPNG(base64PDF string) (*image.Image, error) {

	if len(base64PDF) == 0 {
		return nil, fmt.Errorf("base64 PDF is empty")
	}

	decoded, err := base64.StdEncoding.DecodeString(base64PDF)
	if err != nil {
		return nil, err
	}

	fullPathIn, err := MemoryDataToTmpFile(decoded, "delivrio-label-render", "label.pdf")
	if err != nil {
		return nil, fmt.Errorf("making tmp file: %w", err)
	}
	defer os.RemoveAll(path.Dir(fullPathIn))

	fileIn, err := os.Create(fullPathIn)
	if err != nil {
		return nil, fmt.Errorf("creating input file: %w", err)
	}

	_, err = fileIn.Write(decoded)
	if err != nil {
		return nil, fmt.Errorf("writing input file: %w", err)
	}

	fullPathOut := path.Join(path.Dir(fullPathIn), "label.png")

	cmd := exec.Command(conf.PathPDFium, "render", fullPathIn, fullPathOut)

	cmdOut, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("shelling PDFium: %w: %v: %v", err, string(cmdOut), cmd.String())
	}

	filOut, err := os.Open(fullPathOut)
	if err != nil {
		return nil, fmt.Errorf("open output path: %w", err)
	}

	// load and decode image
	img, _, err := image.Decode(io.Reader(filOut))
	if err != nil {
		return nil, fmt.Errorf("image decode: %w", err)
	}

	return &img, nil
}

func String2Base64(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

// Remember to cleanup with `defer os.RemoveAll(dir)`
func MemoryDataToTmpFile(data []byte, dirPrefix string, fileName string) (string, error) {
	dir, err := os.MkdirTemp("", dirPrefix)
	if err != nil {
		return "", fmt.Errorf("making tmp dir: %w", err)
	}

	fullPathIn := path.Join(dir, fileName)

	fileIn, err := os.Create(fullPathIn)
	if err != nil {
		return "", nil
	}

	_, err = fileIn.Write(data)
	if err != nil {
		return "", fmt.Errorf("writing input file: %w", err)
	}

	return fullPathIn, nil
}

func EncodePNG(img image.Image) ([]byte, error) {
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, img, nil); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func TruncateString(input string, maxLen int) string {
	if len(input) <= maxLen {
		return input
	}
	return input[0:maxLen]
}

func HashPasswordX(password string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func RandomX(length int) string {

	const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	by := make([]byte, length)

	if _, err := rand.Read(by); err != nil {
		panic(err)
	}

	for i, b := range by {
		by[i] = chars[b%byte(len(chars))]
	}

	return string(by)
}

func DeleteImage(p string) (err error) {

	imgPath := path.Join("static", "uploads")

	if strings.HasPrefix(p, imgPath) {
		bucket, err := fileblob.OpenBucket("./uploads", nil)
		if err != nil {
			return err
		}

		err = bucket.Delete(context.Background(), strings.TrimPrefix(p, imgPath))
		if gcerrors.Code(err) == gcerrors.NotFound {
			// Treat not-found as success to allow cleanup
			// in case of error
			return nil
		} else if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("expected path prefix %s, got %s", imgPath, p)
}

func SaveImage(imgBase64 string) (imgPath string, err error) {
	imgType := ""
	if strings.Contains(imgBase64, "data:image/png;base64,") {
		imgType = "png"
	} else if strings.Contains(imgBase64, "data:image/jpg;base64,") || strings.Contains(imgBase64, "data:image/jpeg;base64,") {
		imgType = "jpg"
	} else if imgURL, err := url.Parse(imgBase64); err == nil {
		return imgURL.RequestURI(), nil
	} else {
		return "", errors.New("image type not supported")
	}

	withoutPrefix := regexp.MustCompile("^data:image/[a-z]+;base64,")
	unbased, err := base64.StdEncoding.DecodeString(withoutPrefix.ReplaceAllString(imgBase64, ""))
	if err != nil {
		return "", err
	}

	img, _, err := image.Decode(bytes.NewReader(unbased))
	if err != nil {
		return "", err
	}

	newImage := resize.Resize(160, 0, img, resize.Lanczos3)
	imgName := fmt.Sprintf("%s.%s", uuid.New().String(), imgType)
	imgPath = path.Join("static", "uploads", imgName)

	bucket, err := fileblob.OpenBucket("./uploads", nil)
	if err != nil {
		return "", err
	}

	blob, err := bucket.NewWriter(context.TODO(), imgName, nil)
	if err != nil {
		return "", err
	}
	defer blob.Close()

	if imgType == "png" {
		err = png.Encode(blob, newImage)
		if err != nil {
			return "", err
		}
	} else if imgType == "jpg" {
		err = jpeg.Encode(blob, newImage, nil)
		if err != nil {
			return "", err
		}
	}

	return imgPath, err
}

func SaveFile(fileBase64 string, extension string) (filePath string, err error) {

	withoutPrefix := regexp.MustCompile("^data:application/pdf;base64,")
	unbased, err := base64.StdEncoding.DecodeString(withoutPrefix.ReplaceAllString(fileBase64, ""))
	if err != nil {
		return "", err
	}

	bucket, err := fileblob.OpenBucket("./uploads", nil)
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("%s.%s", uuid.New().String(), extension)
	filePath = path.Join("uploads", fileName)

	blob, err := bucket.NewWriter(context.TODO(), fileName, nil)
	if err != nil {
		return "", err
	}
	defer blob.Close()

	_, err = blob.Write([]byte(unbased))
	if err != nil {
		return "", err
	}

	return filePath, nil

}

// Returns are anonymous, so we default to today
func PickupDate(ctx context.Context) (time.Time, error) {
	cli := ent.FromContext(ctx)
	v := viewer.FromContext(ctx)
	pickupDay := user.PickupDayToday

	currentUser, err := cli.User.Query().
		Where(user.ID(v.MyId())).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return time.Time{}, fmt.Errorf("pickup date: user: %w", err)
	} else if !ent.IsNotFound(err) {
		pickupDay = currentUser.PickupDay
	}

	return labelutils.PickupDayToTime(time.Now(), pickupDay), nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Rollback(tx driver.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

func JoinErrors(allErrors []error, sep string) string {
	if len(allErrors) == 0 {
		return ""
	}

	output := allErrors[0].Error()
	for i, e := range allErrors {
		if i == 0 {
			continue
		}

		output += sep + e.Error()

	}

	return output
}

func SplitPDFPagesToB64(multiPagePDF []byte) ([]string, error) {

	readerInput := bytes.NewReader(multiPagePDF)

	pages, err := api.SplitRaw(readerInput, 1, nil)
	if err != nil {
		return nil, err
	}

	encodedPages := make([]string, 0)

	for _, page := range pages {
		pageData, err := io.ReadAll(page.Reader)
		if err != nil {
			return nil, err
		}
		encodedPages = append(encodedPages, base64.StdEncoding.EncodeToString(pageData))
	}

	return encodedPages, nil

}
