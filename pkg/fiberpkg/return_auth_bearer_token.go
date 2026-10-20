package fiberpkg

import "fmt"

func JoinAuthBearerToken(token string) string {
	return fmt.Sprintf("%s %s", "bearer", token)
}
