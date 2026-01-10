---
name: generateTestsAndBenchmark
description: Generate unit tests and benchmark tests for the selected code
argument-hint: The selected code to generate tests and benchmarks for
---
Given the selected code, generate comprehensive unit test cases and a benchmark test. 
- Ensure tests cover normal cases, edge cases, and special input scenarios.
- For unit tests, verify correctness of return values and side effects (such as global state changes).
- For benchmark tests, measure the performance of the selected function or method under typical usage.
- Reset any global or package-level state before each test to avoid cross-test contamination.
- Output the tests in idiomatic style for the detected language and testing framework.
- Do not include project-specific or file-specific details; generalize for any function or method.
