package fieldjson

type NoteAttributes map[string]string

// Not fully implemented yet in the GQLGen part

/*// UnmarshalGQL implements the graphql.Unmarshaler interface
func (n *NoteAttributes) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case map[string]interface{}:
		name, ok := v["name"].(string)
		if !ok {
			return fmt.Errorf("expected name to be a string")
		}
		value, ok := v["value"].(string)
		if !ok {
			return fmt.Errorf("expected value to be a string")
		}
		n.Name = name
		n.Value = value
		return nil
	default:
		return fmt.Errorf("unexpected type %T", v)
	}
}

// MarshalGQL implements the graphql.Marshaler interface
func (n *NoteAttributes) MarshalGQL(w io.Writer) {
	j, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}
	_, err = io.WriteString(w, string(j))
	if err != nil {
		panic(err)
	}
}*/
