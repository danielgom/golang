# This one goes in .config/fish
if status is-interactive
    starship init fish | source
    fish_add_path /opt/homebrew/bin

    alias ls="ls -laGH"

    alias k=kubectl
    alias current_branch="git branch --show-current"
    alias gp='git push origin $(current_branch)'
    alias gpf='git push origin $(current_branch) --force'
    alias ga='git add .; git reset HEAD CLAUDE.md .junie/guidelines.md 2>/dev/null'
    alias ca='git commit --amend --no-edit --no-verify'
    alias gac='ga && ca'

    # Java
    set -x JAVA_HOME $HOME/Library/Java/JavaVirtualMachines/temurin-21.0.3/Contents/Home
    # set -x GIT_EDITOR nvim

    # Golang
    set -x GOPATH $HOME/go
    set -U fish_user_paths $HOME/go/bin $fish_user_paths
    set -U fish_user_paths /usr/local/protobuf/bin $fish_user_paths

    # Rust
    set -U fish_user_paths $HOME/.cargo/bin $fish_user_paths

end

# Update with user name
alias claude="/Users/<username>/.claude/local/claude"