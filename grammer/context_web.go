package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", authMiddleware(handler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	var currentUser User 

	if v := r.Context().Value("current_user"); v == nil{
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
		return 
	}else{
		u, ok := v.(User)
		if !ok {
			//타입이 User가 아니면 401 에러 리턴
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return 
		}
		currentUser = u  
	}
	fmt.Fprintf(w, "Hi I am %s", currentUser.Name)
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		//사용자의 현재 세션 정보를 기반으로 currentUser 생성
		currentUser, err := getCurrentUser(r)
		if err != nil {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}

		// 2. 기본 컨텍스트에 current_user를 담아 새로운 컨텍스트 생성
		ctx := context.WithValue(r.Context(), "current_user", currentUser)

		// 3. 새로 생성한 컨텍스트 할당한 새로운 `http.Request` 생성
		nextRequest := r.WithContext(ctx)

		// 4. 다음 핸들러 호출
		next(w, nextRequest)
	}
}