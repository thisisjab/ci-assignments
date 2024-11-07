# CI Assignments

This repository contains all assignments for computational intelligence class
(Dr. Koohestani). Included assignments are:

- And Gate with Hebb's rule, single-layer perceptron, and adaline
- X and O character detection with Hebb's rule, single-layer perceptron, and adaline

You can adjust the values for theta, stop condition, learning rate, etc. to re-train the network
through the GUI.

# Where to go

**training:** This package contains all algorithms and training sets used in various
assignments. Please consider that these algorithms are generic and used in all assignments. 

**saved_wights:** This package contains calculated weights plus any useful info. such as theta, learning rate,
count of training vectors after training has finished.

**models:** This package contains all data types used throughout the whole application.

# How to run?

You need to have golang 1.23 installed. Then refer to [fyne installation guide](https://docs.fyne.io/started/)
to install GUI-related dependencies. Then inside root of the project, type in `go mod tidy`. Now you can run
the project with `go run main/main.go`.