package route
import(
	"fmt"
	"ChitChat/internal/handler"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	fmt.Println("User routes")
	router.HandleFunc("/signup", handler.SignUpUser).Methods("POST")
	router.HandleFunc("/login", handler.LoginUser).Methods("POST")
}