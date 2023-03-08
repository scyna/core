#!/bin/bash

read -p "Enter tag name: " tag
read -p "Enter comment: " comment

echo "Creating tag $tag with comment \"$comment\"..."

git tag -a "$tag" -m "$comment"

echo "Pushing tag $tag to remote repository..."

git push origin "$tag"

echo "Tag $tag created and pushed successfully!"
