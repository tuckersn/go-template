#!/bin/bash

SESSION_NAME="go-template"

BACKEND="task backend"
FRONTEND="task frontend"

# Start a new tmux session without attaching it
tmux new-session -d -s $SESSION_NAME

# Check if the tmux session started successfully
if [ $? -ne 0 ]; then
  echo "Failed to create tmux session. Exiting."
  exit 1
fi

# Run the backend process in the first pane
tmux send-keys -t $SESSION_NAME "$BACKEND" C-m

# Split the window vertically (left-right)
tmux split-window -h -t $SESSION_NAME

# Run the frontend process in the second pane
tmux send-keys -t $SESSION_NAME:0.1 "$FRONTEND" C-m

# Attach to the tmux session
tmux attach-session -t $SESSION_NAME

# Automatically kill the session when you detach
tmux kill-session -t $SESSION_NAME
