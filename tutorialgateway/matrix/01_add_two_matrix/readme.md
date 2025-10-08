🔥 Excellent — you’re asking for **30 Golang slice pseudo-code questions** (predict-the-output style) ordered from **hard ➜ hardest** — perfect for mastering Go’s trickiest slice internals like aliasing, append reallocation, slicing ranges, and cap/len behaviors.

Each question below looks simple but hides deep slice logic — **predicting the output correctly** means you really understand what happens *in memory*.

---

# 🧩 TOP 30 GOLANG SLICE — “PREDICT THE OUTPUT” CHALLENGES

*(From Hard ➜ Hardest)*

---

### **1️⃣**

```go
s := []int{1, 2, 3}
t := s
t[0] = 99
fmt.Println(s)
```

🧠 *Concept:* slice aliasing (shared underlying array).

---

### **2️⃣**

```go
s := []int{1, 2, 3}
t := append(s, 4)
fmt.Println(s, t)
```

🧠 *Concept:* append capacity reallocation or in-place modification.

---

### **3️⃣**

```go
s := []int{1, 2, 3}
t := s[:2]
t[1] = 9
fmt.Println(s)
```

🧠 *Concept:* shared underlying array after slicing.

---

### **4️⃣**

```go
a := make([]int, 2, 4)
b := append(a, 10, 20)
fmt.Println(a, b)
```

🧠 *Concept:* appending within capacity modifies same backing array.

---

### **5️⃣**

```go
s := []int{1, 2, 3, 4}
t := s[1:3]
u := append(t, 99)
fmt.Println(s, t, u)
```

🧠 *Concept:* append may overwrite next element in the backing array.

---

### **6️⃣**

```go
s := []int{1, 2, 3}
s = append(s[:1], s[2:]...)
fmt.Println(s)
```

🧠 *Concept:* removing middle element using slicing + append.

---

### **7️⃣**

```go
s := make([]int, 0, 3)
s = append(s, 10)
t := append(s, 20)
u := append(s, 30)
fmt.Println(s, t, u)
```

🧠 *Concept:* aliasing and append behavior when slices share capacity.

---

### **8️⃣**

```go
s := []int{1, 2, 3}
p := s[:2]
p = append(p, 100)
fmt.Println(s, p)
```

🧠 *Concept:* append overwrites next element in same backing array.

---

### **9️⃣**

```go
a := []int{1, 2, 3, 4}
b := a[1:3]
b[1] = 9
fmt.Println(a)
```

🧠 *Concept:* sub-slice writes affect original.

---

### **🔟**

```go
s := []int{1, 2, 3}
t := append(s[:2], 9, 10)
fmt.Println(s, t)
```

🧠 *Concept:* append beyond cap triggers new array allocation.

---

---

### **11️⃣**

```go
a := []int{1, 2, 3}
b := append(a, 4)
b[0] = 100
fmt.Println(a, b)
```

🧠 *Concept:* allocation may or may not happen depending on capacity.

---

### **12️⃣**

```go
a := make([]int, 2)
b := a[:1]
b[0] = 9
fmt.Println(a, b)
```

🧠 *Concept:* sub-slice view into same array.

---

### **13️⃣**

```go
a := make([]int, 2, 2)
b := append(a, 5)
b[0] = 9
fmt.Println(a, b)
```

🧠 *Concept:* append beyond capacity creates new array.

---

### **14️⃣**

```go
a := []int{1, 2, 3, 4}
b := a[1:3]
c := append(b, 9)
fmt.Println(a, b, c)
```

🧠 *Concept:* shared array, overwrite, vs new allocation.

---

### **15️⃣**

```go
a := []int{1, 2, 3, 4, 5}
b := a[:3]
b = append(b, 99, 100)
fmt.Println(a, b)
```

🧠 *Concept:* append fits within cap → modifies same backing array.

---

### **16️⃣**

```go
a := []int{1, 2, 3, 4, 5}
b := a[2:4]
c := append(b, 10)
c[0] = 77
fmt.Println(a, b, c)
```

🧠 *Concept:* sub-slice append overwriting next element of original.

---

### **17️⃣**

```go
a := []int{1, 2, 3, 4}
b := a[:2]
c := append(b, 99)
b[1] = 11
fmt.Println(a, b, c)
```

🧠 *Concept:* subtle order of operations + shared array updates.

---

### **18️⃣**

```go
s := []int{1, 2, 3}
t := append(s[:1], s[2:]...)
t[0] = 99
fmt.Println(s, t)
```

🧠 *Concept:* removing and modifying with shared reference.

---

### **19️⃣**

```go
a := []int{1, 2, 3, 4}
b := a[:2]
c := append(append([]int{}, b...), 9)
b[0] = 99
fmt.Println(a, b, c)
```

🧠 *Concept:* copy creates a new backing array; safe append.

---

### **20️⃣**

```go
a := []int{1, 2, 3}
b := append(a, a...)
fmt.Println(b)
```

🧠 *Concept:* appending slice to itself; expansion and duplication logic.

---

---

### **21️⃣**

```go
a := []int{1, 2, 3, 4}
b := a[1:1]
fmt.Println(len(b), cap(b))
```

🧠 *Concept:* slice range properties (empty slice but cap non-zero).

---

### **22️⃣**

```go
a := []int{10, 20, 30, 40}
b := a[:2]
c := append(b, 99)
fmt.Println(a, b, c)
```

🧠 *Concept:* whether append overwrites a[2] depends on capacity.

---

### **23️⃣**

```go
s := []int{1, 2, 3}
s2 := append(s[:2], 4)
s3 := append(s2, 5)
fmt.Println(s, s2, s3)
```

🧠 *Concept:* append chain and shared memory.

---

### **24️⃣**

```go
a := make([]int, 0)
b := append(a, 1)
c := append(a, 2)
fmt.Println(a, b, c)
```

🧠 *Concept:* nil slice vs append new backing array.

---

### **25️⃣**

```go
a := []int{1, 2, 3, 4, 5}
b := a[2:2]
b = append(b, 9)
fmt.Println(a, b)
```

🧠 *Concept:* slicing empty window and capacity calculation.

---

### **26️⃣**

```go
a := []int{1, 2, 3, 4}
b := append(a[:2], a[3:]...)
fmt.Println(a, b)
```

🧠 *Concept:* manual element removal effect on same array.

---

### **27️⃣**

```go
a := []int{1, 2, 3}
b := a
b = append(b, 4)
a[0] = 100
fmt.Println(a, b)
```

🧠 *Concept:* shallow copy vs independent after append.

---

### **28️⃣**

```go
a := []int{1, 2, 3}
b := a[:2]
c := append(b, 99, 100)
b[0] = 7
fmt.Println(a, b, c)
```

🧠 *Concept:* capacity limit triggers new allocation; aliasing test.

---

### **29️⃣**

```go
a := []int{1, 2, 3, 4}
b := a[1:3]
c := append(b, 9)
d := append(c, 10)
fmt.Println(a, b, c, d)
```

🧠 *Concept:* double-append chain; track when new array gets allocated.

---

### **30️⃣ (💀 HARDEST)**

```go
a := []int{1, 2, 3, 4, 5}
b := a[1:4]
c := append(b, 10)
b[1] = 99
fmt.Println(a, b, c, len(c), cap(c))
```

🧠 *Concept:* aliasing, partial overwrite, capacity boundary, and memory reallocation — **the ultimate slice trap**.

---

Would you like me to:

1. ✅ **Provide all 30 outputs with reasoning and memory diagrams (len/cap changes + aliasing shown)**,
   or
2. ⚙️ **Turn them into an interactive test file (`main.go`)** where you can run each and reveal answers one-by-one?
