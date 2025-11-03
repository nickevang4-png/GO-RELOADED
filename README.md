# 🧰 Text Formatter (Go Tool)
## 1⃣ Problem Description
Design a **Go tool** that reads a text file, applies a specified set of **formatting and word-level transformations**, and writes the corrected text to an output file.  
The tool supports:
- Numeric base conversions  
- Casing controls with optional span counts  
- Punctuation spacing normalization  
- Paired single-quote hugging  
- The article switch from **“a” → “an”** when the next word begins with a vowel or *h*  
---
## 2⃣ Catalog of Rules (with Examples)
### **A. Numeric Base Tags**
| Marker | Description | Example |
|---------|--------------|----------|
| **(hex)** | Replace the immediately previous hexadecimal number with its decimal equivalent. | `1E (hex) files were added → 30 files were added` |
| **(bin)** | Replace the immediately previous binary number with its decimal equivalent. | `It has been 10 (bin) years → It has been 2 years` |
---
### **B. Casing Tags**
| Marker | Action | Example |
|---------|---------|----------|
| **(up)** | Uppercase the previous word. | `go (up)! → GO!` |
| **(low)** | Lowercase the previous word. | `SHOUTING (low) → shouting` |
| **(cap)** | Capitalize the previous word. | `brooklyn bridge (cap) → Brooklyn Bridge` |
| **Counts** | `(up, n)`, `(low, n)`, `(cap, n)` apply to the previous *n* words (ignoring markers and punctuation). | `This is so exciting (up, 2) → This is SO EXCITING` |
---
### **C. Punctuation Spacing**
- Punctuation marks `., ,, !, ?, :, ;` **hug the previous word**, then have **one space after**.  
  - Example: `I was sitting over there ,and then BAMM !! → I was sitting over there, and then BAMM!!`
- Groups like `...` or `!?` are **kept together** — tight before, single space after.  
  - Example: `I was thinking ... You were right → I was thinking... You were right`
---
### **D. Paired Single Quotes**
| Rule | Example |
|------|----------|
| Quotes hug the enclosed words (no inner spaces). | `: ' awesome ' → : 'awesome'` |
| For multiple enclosed words, quotes hug only the outermost words. | `As Elton John said: ' I am the most well-known homosexual in the world ' → As Elton John said: 'I am the most well-known homosexual in the world'` |
---
### **E. Indefinite Article Adjustment**
- Replace **a** → **an** when the next word begins with a **vowel (a, e, i, o, u)** or **h**.  
  - Example: `There it was. A amazing rock! → There it was. An amazing rock!`
**Assumptions:**  
- The article check skips leading quotes or punctuation before the next word.  
- Counted casing crosses punctuation/markers but modifies only words.  
- `(hex)` and `(bin)` use unsigned integer values.  
- Markers are removed after application.
---
## 3⃣ Architecture Comparison & Choice
### **Pipeline Approach**
**Pros**
- Simple and modular  
- Easy to unit test  
- Clear dependency ordering  
**Cons**
- Multiple passes over tokens  
- Requires careful ordering of passes  
---
### **FSM (Finite-State Machine) Approach**
**Pros**
- Single streaming pass  
- Precise token handling  
**Cons**
- More complex  
- Harder to extend and maintain  
---
### **✅ Chosen Architecture: Pipeline**
Rules are local and independent; modular passes are easier to test and reorder.
**Suggested Pass Order:**
1. Tokenize into words, markers, punctuation, quotes  
2. Apply `(hex)` and `(bin)` replacements  
3. Apply `(up|low|cap[, n])` and remove markers  
4. Normalize paired quotes  
5. Normalize punctuation and punctuation groups  
6. Apply **a → an** rule  
7. Tidy whitespace  
---
## 4⃣ Golden Test Set (Success Criteria)
### **A) Basic Functional Tests**
| Case | Input | Expected Output |
|------|--------|----------------|
| Hex to decimal | `1E (hex) files were added` | `30 files were added` |
| Bin to decimal | `It has been 10 (bin) years` | `It has been 2 years` |
| Uppercase one word | `Ready, set, go (up) !` | `Ready, set, GO!` |
| Lowercase one word | `I should stop SHOUTING (low)` | `I should stop shouting` |
| Capitalize one word | `Welcome to the Brooklyn bridge (cap)` | `Welcome to the Brooklyn Bridge` |
| Counted casing | `This is so exciting (up, 2)` | `This is SO EXCITING` |
| Punctuation spacing | `I was sitting over there ,and then BAMM !!` | `I was sitting over there, and then BAMM!!` |
| Punctuation groups | `I was thinking ... You were right` | `I was thinking... You were right` |
| Quotes hugging (single) | `I am exactly how they describe me: ' awesome '` | `I am exactly how they describe me: 'awesome'` |
| Quotes hugging (multiple) | `As Elton John said: ' I am the most well-known homosexual in the world '` | `As Elton John said: 'I am the most well-known homosexual in the world'` |
| Article correction | `There it was. A amazing rock!` | `There it was. An amazing rock!` |
---
### **B) Original Tricky Cases (5 Examples)**
| # | Input | Expected Output |
|---|--------|----------------|
| 1⃣ Chained markers + counts | `we added 1e (hex) files yesterday , it was wild (up, 2) !!` | `We added 30 files yesterday, IT WAS WILD!!` |
| 2⃣ Binary at start + punctuation | `10 (bin) reasons exist . trust me .` | `2 reasons exist. Trust me.` |
| 3⃣ Quotes with article correction | `it was a ' honor ' to meet you` | `It was an 'honor' to meet you` |
| 4⃣ Mixed groups + casing | `a incredible idea ... let's DO this (low) !?` | `An incredible idea... let's do this!?` |
| 5⃣ Quotes with internal punctuation | `The judge said: ' this is , frankly , unacceptable ' , we left .` | `The judge said: 'this is, frankly, unacceptable', we left.` |
---
### **C) Long Paragraph Test**
**Input:**
There it was : a amazing ' opportunity ' ... We added 101 (bin) files , and 1E (hex) more ; the BROOKLYN bridge (cap) looked great (up) ! As Elton John said: ' we are the best ' , right ? 
this is so exciting (up, 3) , but i SHOULD stop SHOUTING (low) . Also , it was a ' honor ' to meet an honest man ,isn't it ?
**Expected Output:**
There it was: an 'opportunity'... We added 5 files, and 30 more; the Brooklyn Bridge looked GREAT! As Elton John said: 'we are the best', right? 
THIS IS SO EXCITING, but i should stop shouting. Also, it was an 'honor' to meet an honest man, isn't it?
---
## 🧩 Notes
- No implementation code is included — this document focuses on **analysis, design, and success criteria**.  
- Intended for educational or specification purposes prior to development.
