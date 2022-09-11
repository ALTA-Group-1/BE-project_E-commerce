package delivery

import "project/e-commerce/features/user"

type UserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email,omitempty"`
	Address string `json:"address,omitempty"`
}

func fromCore(data user.Core) UserResponse {
	return UserResponse{
		ID:      data.ID,
		Name:    data.Name,
		Email:   data.Email,
		Address: data.Address,
	}
}
