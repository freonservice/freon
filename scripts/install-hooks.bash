#!/usr/bin/env bash

GIT_DIR=$(git rev-parse --git-dir)

echo "Installing hooks..."
rm $GIT_DIR/hooks/pre-commit
# this command creates symlink to our pre-commit script
ln -s ../../scripts/pre-commit.bash $GIT_DIR/hooks/pre-commit
chmod +x $GIT_DIR/hooks/pre-commit

echo "Done!"