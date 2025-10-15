package dto

// ImageUpdate represents image modification options for a product.
type ImageUpdate struct {
	Add     []string `json:"add,omitempty"`    
	Delete  []string `json:"delete,omitempty"` 
	URLs    []string `json:"urls,omitempty"`   
	Replace bool     `json:"replace,omitempty"` 
}
