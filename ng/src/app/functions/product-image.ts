type ProductImage = {productImage?: {url: string;}[]};
export type ProductVariantImages = ProductImage & {product: ProductImage};

export function defaultProductImg(input: ProductVariantImages): string {
	if (!!input.productImage && input.productImage.length > 0) {
		return input.productImage[0]?.url || '';
	}

	if (!!input.product?.productImage && input.product?.productImage.length > 0) {
		return input.product.productImage[0]?.url || '';
	}

	return "";
}
