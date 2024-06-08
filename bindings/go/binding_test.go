package tree_sitter_bitbake_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-bitbake"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_bitbake.Language())
	if language == nil {
		t.Errorf("Error loading Bitbake grammar")
	}
}
