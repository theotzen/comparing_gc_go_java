# Garbage Collection Analysis Tool

## Overview

This tool accompanies the technical article "Garbage Collectors: Brief History and Comparison of Go and Java" and provides a practical setup for analyzing and comparing the garbage collection (GC) mechanisms in Go and Java. It leverages containerization, Prometheus, and Grafana to monitor and visualize GC performance metrics under simulated memory churn conditions.

## Features

- Generates large linked lists to simulate memory churn.
- Exposes GC metrics for both Go and Java applications.
- Integrates with Prometheus for metric collection.
- Provides Grafana dashboards for real-time visualization and comparison.

## Metrics Exposed

- Total number of GC operations
- Cumulative GC time
- Used allocated heap bytes
- Target heap size of the next GC cycle (Go-specific)
- Total system obtained bytes (Go-specific)

## Usage

The applications are designed to run indefinitely, generating and destroying linked lists to create a continuous load for the garbage collectors. Prometheus will scrape the exposed metrics at a predefined interval, and Grafana will display the data on the dashboards.
