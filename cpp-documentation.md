# C++ Documentation

## Overview

This document provides comprehensive guidance for C++ development, covering fundamental concepts, best practices, and advanced features.

## Table of Contents

1. [Getting Started](#getting-started)
2. [Basic Syntax](#basic-syntax)
3. [Data Types](#data-types)
4. [Classes and Objects](#classes-and-objects)
5. [Memory Management](#memory-management)
6. [Templates](#templates)
7. [Standard Library](#standard-library)
8. [Best Practices](#best-practices)
9. [Build System](#build-system)
10. [Testing](#testing)

## Getting Started

### Prerequisites

- C++ compiler (GCC, Clang, or MSVC)
- Build system (CMake recommended)
- Text editor or IDE

### Hello World Example

```cpp
#include <iostream>

int main() {
    std::cout << "Hello, World!" << std::endl;
    return 0;
}
```

## Basic Syntax

### Variables and Constants

```cpp
// Variable declaration
int age = 25;
double price = 19.99;
std::string name = "John";

// Constants
const int MAX_SIZE = 100;
constexpr double PI = 3.14159;
```

### Control Structures

```cpp
// If-else statement
if (age >= 18) {
    std::cout << "Adult" << std::endl;
} else {
    std::cout << "Minor" << std::endl;
}

// For loop
for (int i = 0; i < 10; ++i) {
    std::cout << i << " ";
}

// Range-based for loop (C++11)
std::vector<int> numbers = {1, 2, 3, 4, 5};
for (const auto& num : numbers) {
    std::cout << num << " ";
}
```

## Data Types

### Fundamental Types

- `int`: Integer numbers
- `double`: Floating-point numbers
- `char`: Single characters
- `bool`: Boolean values (true/false)
- `void`: No value

### User-Defined Types

```cpp
// Struct
struct Point {
    double x, y;
};

// Enum
enum class Color { RED, GREEN, BLUE };

// Union
union Data {
    int i;
    float f;
    char str[20];
};
```

## Classes and Objects

### Basic Class Definition

```cpp
class Rectangle {
private:
    double width, height;

public:
    // Constructor
    Rectangle(double w, double h) : width(w), height(h) {}
    
    // Member functions
    double area() const {
        return width * height;
    }
    
    double perimeter() const {
        return 2 * (width + height);
    }
    
    // Getter and setter
    double getWidth() const { return width; }
    void setWidth(double w) { width = w; }
};
```

### Inheritance

```cpp
class Shape {
public:
    virtual double area() const = 0;  // Pure virtual function
    virtual ~Shape() = default;       // Virtual destructor
};

class Circle : public Shape {
private:
    double radius;

public:
    Circle(double r) : radius(r) {}
    
    double area() const override {
        return 3.14159 * radius * radius;
    }
};
```

## Memory Management

### Stack vs Heap

```cpp
// Stack allocation
int stackVar = 42;
Rectangle rect(10.0, 5.0);

// Heap allocation (raw pointers - avoid in modern C++)
int* heapVar = new int(42);
delete heapVar;  // Manual cleanup required

// Modern C++ - Smart pointers
std::unique_ptr<int> smartPtr = std::make_unique<int>(42);
std::shared_ptr<Rectangle> sharedRect = std::make_shared<Rectangle>(10.0, 5.0);
```

### RAII (Resource Acquisition Is Initialization)

```cpp
class FileHandler {
private:
    std::FILE* file;

public:
    FileHandler(const std::string& filename) {
        file = std::fopen(filename.c_str(), "r");
        if (!file) {
            throw std::runtime_error("Failed to open file");
        }
    }
    
    ~FileHandler() {
        if (file) {
            std::fclose(file);
        }
    }
    
    // Disable copy constructor and assignment
    FileHandler(const FileHandler&) = delete;
    FileHandler& operator=(const FileHandler&) = delete;
};
```

## Templates

### Function Templates

```cpp
template<typename T>
T maximum(T a, T b) {
    return (a > b) ? a : b;
}

// Usage
int maxInt = maximum(10, 20);
double maxDouble = maximum(3.14, 2.71);
```

### Class Templates

```cpp
template<typename T>
class Stack {
private:
    std::vector<T> elements;

public:
    void push(const T& element) {
        elements.push_back(element);
    }
    
    T pop() {
        if (elements.empty()) {
            throw std::runtime_error("Stack is empty");
        }
        T top = elements.back();
        elements.pop_back();
        return top;
    }
    
    bool empty() const {
        return elements.empty();
    }
};
```

## Standard Library

### Containers

```cpp
#include <vector>
#include <map>
#include <set>
#include <string>

// Vector
std::vector<int> numbers = {1, 2, 3, 4, 5};
numbers.push_back(6);

// Map
std::map<std::string, int> ages;
ages["Alice"] = 30;
ages["Bob"] = 25;

// Set
std::set<int> uniqueNumbers = {3, 1, 4, 1, 5, 9};
```

### Algorithms

```cpp
#include <algorithm>
#include <numeric>

std::vector<int> data = {5, 2, 8, 1, 9};

// Sort
std::sort(data.begin(), data.end());

// Find
auto it = std::find(data.begin(), data.end(), 8);
if (it != data.end()) {
    std::cout << "Found 8 at position " << (it - data.begin()) << std::endl;
}

// Accumulate
int sum = std::accumulate(data.begin(), data.end(), 0);
```

## Best Practices

### Code Style

1. Use meaningful variable and function names
2. Follow consistent naming conventions (camelCase or snake_case)
3. Keep functions small and focused
4. Use const whenever possible
5. Prefer initialization over assignment

### Modern C++ Features

```cpp
// Auto keyword (C++11)
auto result = someComplexFunction();

// Range-based for loops (C++11)
for (const auto& item : container) {
    // Process item
}

// Lambda functions (C++11)
auto lambda = [](int x, int y) { return x + y; };

// Smart pointers (C++11)
std::unique_ptr<Object> ptr = std::make_unique<Object>();

// Move semantics (C++11)
std::vector<int> source = {1, 2, 3};
std::vector<int> dest = std::move(source);
```

### Error Handling

```cpp
#include <stdexcept>

class Calculator {
public:
    double divide(double a, double b) {
        if (b == 0.0) {
            throw std::invalid_argument("Division by zero");
        }
        return a / b;
    }
};

// Usage with exception handling
try {
    Calculator calc;
    double result = calc.divide(10.0, 0.0);
} catch (const std::exception& e) {
    std::cerr << "Error: " << e.what() << std::endl;
}
```

## Build System

### CMake Example

```cmake
cmake_minimum_required(VERSION 3.12)
project(MyProject)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Add executable
add_executable(myapp
    src/main.cpp
    src/calculator.cpp
)

# Add include directories
target_include_directories(myapp PRIVATE include)

# Link libraries
find_package(Threads REQUIRED)
target_link_libraries(myapp Threads::Threads)
```

### Compilation Commands

```bash
# Using g++
g++ -std=c++17 -Wall -Wextra -O2 -o myapp main.cpp

# Using CMake
mkdir build
cd build
cmake ..
make
```

## Testing

### Unit Testing with Google Test

```cpp
#include <gtest/gtest.h>
#include "calculator.h"

class CalculatorTest : public ::testing::Test {
protected:
    Calculator calc;
};

TEST_F(CalculatorTest, Addition) {
    EXPECT_EQ(calc.add(2, 3), 5);
    EXPECT_EQ(calc.add(-1, 1), 0);
}

TEST_F(CalculatorTest, Division) {
    EXPECT_DOUBLE_EQ(calc.divide(10, 2), 5.0);
    EXPECT_THROW(calc.divide(10, 0), std::invalid_argument);
}

int main(int argc, char** argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
```

## Conclusion

This documentation covers the essential aspects of C++ development. For more advanced topics, refer to the official C++ documentation and standards. Remember to always follow modern C++ practices and leverage the rich standard library for efficient and safe code development.

## References

- [C++ Reference](https://en.cppreference.com/)
- [ISO C++ Guidelines](https://isocpp.github.io/CppCoreGuidelines/)
- [Modern C++ Features](https://github.com/AnthonyCalandra/modern-cpp-features)