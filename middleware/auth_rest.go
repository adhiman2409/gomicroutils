package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/adhiman2409/gomicroutils/grpcclient"
	"github.com/gorilla/mux"
)

func RequestAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ai := grpcclient.AuthInfo{
			Authorised:  false,
			Tenant:      "",
			Domain:      "",
			Department:  "",
			Name:        "",
			EmailId:     "",
			PhoneNumber: "",
			Role:        "anonymous",
		}
		byteArray, _ := json.Marshal(ai)

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			ctx := context.WithValue(r.Context(), "claims", string(byteArray))
			fmt.Println("Token length wrong or Token missing")
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		jwtToken := authHeader[1]
		api := mux.CurrentRoute(r).GetName()
		claims, err := grpcclient.GetAuthClient().Verify(jwtToken, api)
		if err != nil {
			fmt.Println("Token Error: " + err.Error())
			ctx := context.WithValue(r.Context(), "claims", string(byteArray))
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		ai.Authorised = claims.Authorised
		ai.Name = claims.Name
		ai.EmailId = claims.EmailId
		ai.PhoneNumber = claims.PhoneNumber
		ai.Role = claims.Role
		ai.Department = claims.Department
		ai.Domain = claims.Domain
		ai.Tenant = claims.Tenant

		byteArray, _ = json.Marshal(ai)

		ctx := context.WithValue(r.Context(), "claims", string(byteArray))

		fmt.Printf("%+v User Authorized for api %s\n", ai, api)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
