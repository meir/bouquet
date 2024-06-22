package client

import (
	"bytes"
	"fmt"

	"github.com/bluekeyes/go-gitdiff/gitdiff"
)

// applyPatch will apply a patch to the asar file
// for example:
//
// Index: app/index.js
// --- app/index.js
// +++ app/index.js
// @@ -90,7 +90,9 @@
//
//		  mainScreen = require('./mainScreen');
//		  mainScreen.init();
//	  + // Bouquet injection
//	  + require('../bouquet').init();
//	    const {
//	      getWindow: getPopoutWindowByKey,
//	      getAllWindows: getAllPopoutWindows,
//	      setNewWindowEvent
//	    } = require('./popoutWindows');
//
// This patch will be seen by the injector by its .patch extension and will then run this method
// This method will then parse the patch and run those patches on the files it can find using the ASAR file.
// Once the patch is applied, it will save the new content in the file header
func (c *Client) applyPatch(p []byte) error {
	diffs, preamble, err := gitdiff.Parse(bytes.NewReader(p))
	if err != nil {
		return fmt.Errorf("failed to parse patch: %w", err)
	}

	fmt.Println(preamble)

	for _, diff := range diffs {
		header := c.asarFile.Header.Get(diff.OldName)
		if header == nil {
			fmt.Printf("no header for %s\n", diff.OldName)
			continue
		}

		fmt.Println("Patching", header.Name())

		var output bytes.Buffer
		err := gitdiff.Apply(&output, bytes.NewReader(header.Content()), diff)
		if err != nil {
			return fmt.Errorf("failed to apply patch for '%s': %w", diff.OldName, err)
		}

		header.SetContent(output.Bytes())
	}

	return nil
}
