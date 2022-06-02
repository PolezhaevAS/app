package token

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go/v4"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// well known HMAC singing methods.
const (
	HS256 = "HS256"
	HS384 = "HS384"
	HS512 = "HS512"
)

// The Config represents configurations for the JWT Source.
type Config struct {
	// Secret to sign tokens.
	Secret string `yaml:"secret" mapstructure:"secret"`
	// Algorithm represents preffered siging method.
	// Supported signing.
	Algorithm string `yaml:"algorithm" mapstructure:"algorithm"`
	// AcceptAlgorithms represents the list of supported algorithms with which a JWSN token
	// can be accepted. The Algorithm field is included in the list
	// automatically. It used to change singing algorithm and let existing
	// tokens work until they expired.
	AcceptAlgorithms []string `yaml:"accept_algorithms" mapstructure:"accept_algorithms"`
}

// NewConfig returns the new default configurations with empty secret.
func NewConfig() *Config {
	return &Config{
		Secret:           "",
		Algorithm:        HS256,
		AcceptAlgorithms: []string{},
	}
}

// The Source used to initiate, parse and verfiy jwt tokens.
type Source struct {
	// app secret key
	secret []byte
	alg    jwt.SigningMethod
	accept []string
}

// New creates a new jwn instance in give config.
func New(conf *Config) (
	s *Source, err error) {
	var alg jwt.SigningMethod
	if alg, err = getAlgorithm(conf.Algorithm); err != nil {
		return nil, err
	}

	// don't allow empty secret for security reasons.
	if conf.Secret == "" {
		return nil, errors.New("empty secret key")
	}

	return &Source{
		secret: []byte(conf.Secret),
		alg:    alg,
		accept: append([]string{conf.Algorithm}, conf.AcceptAlgorithms...),
	}, nil
}

func (s *Source) Sign(cl *Claims) (
	string, error) {
	var jwt = jwt.NewWithClaims(s.alg, cl)
	return jwt.SignedString([]byte(s.secret))
}

func getAlgorithm(name string) (
	alg jwt.SigningMethod, err error) {
	switch name {
	case HS256:
		alg = jwt.SigningMethodHS256
	case HS384:
		alg = jwt.SigningMethodHS384
	case HS512:
		alg = jwt.SigningMethodHS512
	default:
		err = fmt.Errorf("unknown singing algorithm: %q", name)
	}
	return
}

// Parse and verify and validate given token string.
func (s *Source) Parse(tok string) (
	cl *Claims, err error) {

	var token *jwt.Token
	cl = new(Claims)
	if token, err = jwt.ParseWithClaims(tok, cl, s.keyFunc); err != nil {
		return nil, s.checkErr(err)
	}

	if !s.isAcceptable(token.Method.Alg()) {
		return nil, status.Error(codes.Unauthenticated, "invalid authorization token: unsupported singing algorithm")
	}

	return cl, nil
}

func (s *Source) keyFunc(*jwt.Token) (
	interface{}, error) {

	return []byte(s.secret), nil
}

func (s *Source) checkErr(err error) error {
	// TODO: check error condition if needed.
	st, ok := status.FromError(err)
	if st != nil {
		if ok {
			return st.Err()
		}
	}
	return status.Error(codes.Unauthenticated, "invalid authorization token")
}

func (s *Source) isAcceptable(alg string) (
	ok bool) {
	for _, name := range s.accept {
		if alg == name {
			return true
		}
	}
	return // no in the list
}
