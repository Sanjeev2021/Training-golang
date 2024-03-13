// // Online C compiler to run C program online
// #include <iostream>

// using namespace std;

// class Stack {
// public:
//     int *arr;
//     int top;
//     int size;

//     Stack(int size) {
//         this->size = size;
//         arr = new int[size];
//         top = -1;
//     }

//     void push(int element) {
//         if (top < size - 1) {
//             top++;
//             arr[top] = element;
//         } else {
//             cout << "Stack overflow" << endl;
//         }
//     }

//     void pop() {
//         if (top >= 0) {
//             top--;
//         } else {
//             cout << "Stack underflow" << endl;
//         }
//     }

//     int peek() {
//         if (top >= 0 && top < size)
//             return arr[top];
//         else {
//             cout << "Stack is empty" << endl;
//             return -1;
//         }
//     }

//     bool isEmpty() {
//         return top == -1;
//     }
// };

// int main() {
//     // Write C++ code here
//     Stack st(5);
//     st.push(22);
//     st.push(43);

//     cout << st.peek() << endl;
//     st.pop();

//     if (st.isEmpty()) {
//         cout << "Stack is empty" << endl;
//     } else {
//         cout << "Stack is not empty" << endl;
//     }

//     return 0;
// }
