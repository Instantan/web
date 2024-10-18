package generate

import (
	"net/http"
	"testing"
)

func TestGenerateFunctionName(t *testing.T) {
	tests := []struct {
		method       string
		route        string
		expectedName string
	}{
		{
			method:       http.MethodGet,
			route:        "/group/{name}/items/{name}",
			expectedName: "GetGroupItemsByNameAndName2",
		},
		{
			method:       http.MethodPost,
			route:        "/users/{userID}/posts",
			expectedName: "CreateUsersPostsByUserID",
		},
		{
			method:       http.MethodDelete,
			route:        "/products/{productID}/reviews/{reviewID}",
			expectedName: "DeleteProductsReviewsByProductIDAndReviewID",
		},
		{
			method:       http.MethodPut,
			route:        "/articles/{id}",
			expectedName: "UpdateArticlesById",
		},
		{
			method:       http.MethodPatch,
			route:        "/articles/{id}/comments/{commentId}",
			expectedName: "UpdateArticlesCommentsByIdAndCommentId",
		},
		{
			method:       http.MethodGet,
			route:        "/",
			expectedName: "Get",
		},
		{
			method:       http.MethodPost,
			route:        "/login",
			expectedName: "CreateLogin",
		},
		{
			method:       http.MethodGet,
			route:        "/users",
			expectedName: "GetUsers",
		},
		{
			method:       http.MethodGet,
			route:        "/users/{userID}/profile",
			expectedName: "GetUsersProfileByUserID",
		},
		{
			method:       http.MethodGet,
			route:        "/users/{userID}/posts/{postID}",
			expectedName: "GetUsersPostsByUserIDAndPostID",
		},
		{
			method:       "UNKNOWN",
			route:        "/unknown/method",
			expectedName: "UnknownUnknownMethod",
		},
		{
			method:       http.MethodGet,
			route:        "/api/v1/resources/{resourceID}/data/{dataID}",
			expectedName: "GetApiV1ResourcesDataByResourceIDAndDataID",
		},
		{
			method:       http.MethodDelete,
			route:        "/items/{id}/",
			expectedName: "DeleteItemsById",
		},
		{
			method:       http.MethodOptions,
			route:        "/options/test",
			expectedName: "OptionsOptionsTest",
		},
		{
			method:       http.MethodHead,
			route:        "/head/test",
			expectedName: "HeadHeadTest",
		},
	}

	for _, tt := range tests {
		t.Run(tt.method+" "+tt.route, func(t *testing.T) {
			actualName := generateFunctionName(tt.method, tt.route)
			t.Log(actualName)
			if actualName != tt.expectedName {
				t.Errorf("generateFunctionName(%q, %q) = %q; want %q", tt.method, tt.route, actualName, tt.expectedName)
			}
		})
	}
}
