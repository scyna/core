#!/bin/bash
#sh publish.sh v1.0.0
git tag $1
git push origin $1
GOPROXY=proxy.golang.org go list -m github.com/scyna/core@$1

# read -p "Enter tag name: " tag
# read -p "Enter comment: " comment

# echo "Creating tag $tag with comment \"$comment\"..."

# git tag -a "$tag" -m "$comment"

# echo "Pushing tag $tag to remote repository..."

# git push origin "$tag"

# echo "Tag $tag created and pushed successfully!"
