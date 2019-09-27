module github.com/ehlxr/jwt

go 1.13

require (
	github.com/atotto/clipboard v0.1.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/manifoldco/promptui v0.0.0-20180308161052-c0c0d3afc6a0
	github.com/mitchellh/go-homedir v0.0.0-20161203194507-b8bc1bf76747
	github.com/spf13/cobra v0.0.2
	github.com/spf13/viper v1.0.2
	jwt v0.0.0-00010101000000-000000000000
)

replace jwt => ./.
