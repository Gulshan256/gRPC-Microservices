package main

import "github.com/99designs/gqlgen/graphql"

type Server struct {
	// accountClient *account.client
	// orderClient *order.client
	// catalogClient *catalog.client
}

func NewGraphQLServer(accounturl, catalogurl, orderUrl string) (*Server, error) {
	// 	accoountClient, err := account.NewClient(accounturl)
	// 	if err != nil {
	// 		accoountClient.close()
	// 		return nil, err
	// 	}

	// 	catalogClient, err := catalog.NewClient(catalogurl)
	// 	if err != nil {
	// 		catalogClient.close()
	// 		return nil, err
	// 	}

	// 	orderClient, err := account.NewClient(orderUrl)
	// 	if err != nil {
	// 		orderClient.close()
	// 		return nil, err
	// 	}

	return &Server{
		// 		accoountClient,
		//  	catalogClient,
		//  	orderClient),
	}, nil
}

// func (s *Server) Mutation() MutationResolver {
// 	return &mutationResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Query() queryResolver {
// 	return &queryResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Account() accountResolver {
// 	return &accountResolver{
// 		server: s,
// 	}
// }

func (s *Server) ToExcutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(
		Config{
			Resolvers: s,
		},
	)
}
