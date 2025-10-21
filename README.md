# Text Editor Project - Analysis Document

## 1. Περιγραφή του Προβλήματος

Το πρόβλημα αφορά τη δημιουργία ενός εργαλείου επεξεργασίας κειμένου που διαβάζει ένα αρχείο εισόδου, εφαρμόζει μια σειρά από μετασχηματισμούς και κανόνες μορφοποίησης, και γράφει το αποτέλεσμα σε αρχείο εξόδου.

Το πρόγραμμα πρέπει να αναγνωρίζει ειδικούς modifiers μέσα στο κείμενο (π.χ. `(hex)`, `(up)`, `(cap)`) και να εφαρμόζει τους αντίστοιχους μετασχηματισμούς στις προηγούμενες λέξεις. Επιπλέον, πρέπει να διορθώνει αυτόματα τη στίξη, τα κενά γύρω από σημεία στίξης, και να χειρίζεται ειδικές περιπτώσεις όπως τα εισαγωγικά και τη διόρθωση άρθρων (a/an).

Η πρόκληση είναι να γίνει αυτό αποδοτικά, με σωστή διαχείριση του context και με τρόπο που να είναι εύκολα επεκτάσιμος και maintainable.

![FSM Diagram](https://www.mermaidchart.com/play?utm_source=mermaid_live_editor&utm_medium=toggle#pako:eNp1VP1v2jAQ_Ve8TCqbJgpdt46yqV0gASGtoYWgahpTlSZHsQh25jj9WNX_feePhBTGT7HP7907n9_l2Yl5Ak7XaTabcxZztqB33TkjJI2eeCG7BNKV2solrKFLkkjgVmMXKX-Il5GQJOwpBCFTibt3v-bONHQn4dz5_Z40m2dkxKjEoPrQKKV_gfSKxQIEOVAMCQg0fIXQjAlEyfPcUR8SwCMq8BWwufNicDrebJJrLhKNVwtUmPiuNwqGN9fjibeVuqRMsygGzdErJF0KHkOem_0O_qrgEkhDM_QaGVezcejfhOPh8Ie_R-ayYLEsIkk501S9V2KzoB_O3HA0DvYw_fFAM3yWqE76gUfwooJiFeNCZoVUbTWU6vr9JcSrC65aNsoJLuiCgjj_divOlvDYuqWsVWQtfK5WHGVVF0uWUv0Juc7kZln6hDFU1ksSiojlCy7W-i4qo-TkUsA95UWuK2jl1R3qGQO-Kc0NsDK9Im7LZWRSpFCVUUpquDEGqs-yBLtj968F3EDlb0QN8oHc8wdIW0vNHdBHN_j-jMa9hbRLlGLE7kAjsehGxBpzh7zsZMJKY4NkAAnUCzFYnXgnaj2swuZhjlTTkkRpVS9l58KaDg2e0wTIH2Uk0_CAlz6crmi2ZcKKF6gHSM0LJUnJUGIGsitp81bTZHttyVthY3LUGVCRS-v2cQbsClUucOItQs93JVGxpoC_jcTS-inPYYuHXlbOGWgbadNsPKNltuoxOWqt_bi3tXqwNHIoeJGhz2qjp0R1WI_C4eEh4YK8Oa-sp8_q9u_zNQ6L6q1dkYNonX0lpnKVBM1CmQTB8Dly1crNTaps1vpTyu5Slcve2-xVkqmkOAp4mawcpAfsycblVntz_-P91tJJ61DbVmPJrcbabv43erwVLccpX0YZ_vgFxNK2LZdPqTUDWdA07b49bfv-abt-ah7ennYGvdNO_RQNYc8Gg95J_6h-Vv1ESoD_qfe5Dqh-GAbQ-dL3_d5OBpxYc-55btvz6udmnm36tn_S6ddPzTO-Fnde_gG-lzj-)

## 2. Σύγκριση Αρχιτεκτονικών

### Pipeline Architecture

**Πώς λειτουργεί:**
Το κείμενο περνάει από μια σειρά ανεξάρτητων φίλτρων/stages:
1. Stage 1: Hex/Bin conversions
2. Stage 2: Case transformations
3. Stage 3: Punctuation formatting
4. Stage 4: Quote handling
5. Stage 5: A/An correction

**Πλεονεκτήματα:**
- Απλή και γραμμική λογική
- Κάθε stage είναι ανεξάρτητο και testable
- Εύκολη προσθήκη νέων stages
- Separation of concerns

**Μειονεκτήματα:**
- Χρειάζεται πολλαπλά passes στο κείμενο (5+ φορές)
- Κάθε stage δημιουργεί intermediate string
- Μεγάλη χρήση μνήμης για μεγάλα αρχεία
- Δύσκολο να διατηρήσεις context μεταξύ stages
- Πιο αργό για μεγάλα inputs

### FSM (Finite State Machine) Architecture 

**Πώς λειτουργεί:**
Το πρόγραμμα βρίσκεται πάντα σε ένα συγκεκριμένο "state" και διαβάζει το input character-by-character ή token-by-token. Ανάλογα με το τρέχον state και το επόμενο input, αλλάζει state και εκτελεί actions.

**Core States:**
- `READING_WORD`: Συλλογή χαρακτήρων μιας λέξης
- `WORD_COMPLETE`: Λέξη ολοκληρώθηκε, έλεγχος για modifiers
- `IN_QUOTES`: Tracking αν είμαστε μέσα σε εισαγωγικά
- `READING_MODIFIER`: Ανάγνωση (hex), (up), κλπ
- `HANDLE_PUNCTUATION`: Εφαρμογή κανόνων στίξης

**Πλεονεκτήματα:**
-  **Single pass** - διαβάζει το κείμενο μόνο μία φορά
-  **Memory efficient** - δεν δημιουργεί intermediate copies
-  **Context awareness** - το state machine "θυμάται" που βρίσκεται (μέσα σε quotes, μετά από modifier, κλπ)
-  **Natural fit** - το text parsing είναι κλασικό FSM πρόβλημα
-  **Faster execution** - O(n) complexity με minimal overhead
-  **Easier debugging** - ξέρεις ακριβώς σε ποιο state βρίσκεσαι
-  **Scalability** - handle μεγάλα αρχεία αποδοτικά

**Μειονεκτήματα:**
- Πιο πολύπλοκος αρχικός σχεδιασμός
- Χρειάζεται προσεκτική σκέψη για τα state transitions
- Harder to modify μετά την αρχική υλοποίηση

### Προσωπική Επιλογή: FSM 

**Γιατί FSM:**

1. **Performance**: Για text processing, το single-pass FSM είναι αντικειμενικά πιο γρήγορο
2. **Memory**: Σε μεγάλα αρχεία, η διαφορά στη μνήμη είναι σημαντική
3. **Natural Design**: Το parsing κειμένου είναι φυσικό FSM πρόβλημα - προσομοιάζει πώς σκέφτεται κανείς όταν διαβάζει
4. **Context Handling**: Πολλοί κανόνες εξαρτώνται από context (είμαι μέσα σε quotes? μόλις είδα modifier?). Το FSM το χειρίζεται φυσικά
5. **Industry Standard**: Compilers, parsers, lexers - όλα χρησιμοποιούν FSM
6. **Learning Value**: Πιο εκπαιδευτικό και επαγγελματικό approach

Το **tradeoff** της πολυπλοκότητας αξίζει για τα benefits που προσφέρει το FSM σε αυτό το project.

---

## 3. Καταγραφή των Κανόνων

### 3.1 Αριθμητικοί Μετασχηματισμοί

#### Κανόνας: `(hex)` - Hexadecimal σε Decimal
- **Περιγραφή**: Μετατρέπει την προηγούμενη λέξη από hex σε decimal
- **Παράδειγμα**: `"1E (hex) files"` → `"30 files"`
- **Παράδειγμα**: `"FF (hex) is max"` → `"255 is max"`

#### Κανόνας: `(bin)` - Binary σε Decimal
- **Περιγραφή**: Μετατρέπει την προηγούμενη λέξη από binary σε decimal
- **Παράδειγμα**: `"10 (bin) years"` → `"2 years"`
- **Παράδειγμα**: `"1010 (bin) equals"` → `"10 equals"`

### 3.2 Μετασχηματισμοί Κεφαλαίων/Πεζών

#### Κανόνας: `(up)` - Uppercase
- **Περιγραφή**: Μετατρέπει την προηγούμενη λέξη σε ΚΕΦΑΛΑΙΑ
- **Παράδειγμα**: `"go (up) now"` → `"GO now"`

#### Κανόνας: `(low)` - Lowercase
- **Περιγραφή**: Μετατρέπει την προηγούμενη λέξη σε πεζά
- **Παράδειγμα**: `"SHOUTING (low)"` → `"shouting"`

#### Κανόνας: `(cap)` - Capitalize
- **Περιγραφή**: Κάνει κεφαλαίο μόνο το πρώτο γράμμα
- **Παράδειγμα**: `"bridge (cap)"` → `"Bridge"`

### 3.3 Μετασχηματισμοί με Αριθμό

#### Κανόνας: `(up, N)` - Uppercase N λέξεων
- **Περιγραφή**: Μετατρέπει τις N προηγούμενες λέξεις σε ΚΕΦΑΛΑΙΑ
- **Παράδειγμα**: `"so exciting (up, 2)"` → `"SO EXCITING"`

#### Κανόνας: `(low, N)` - Lowercase N λέξεων
- **Περιγραφή**: Μετατρέπει τις N προηγούμενες λέξεις σε πεζά
- **Παράδειγμα**: `"IT WAS THE (low, 3)"` → `"it was the"`

#### Κανόνας: `(cap, N)` - Capitalize N λέξεων
- **Περιγραφή**: Κάνει capitalize τις N προηγούμενες λέξεις
- **Παράδειγμα**: `"age of foolishness (cap, 3)"` → `"Age Of Foolishness"`

### 3.4 Κανόνες Στίξης

#### Κανόνας: Βασική Στίξη (. , ! ? : ;)
- **Περιγραφή**: Κολλάει στην προηγούμενη λέξη, κενό μετά
- **Παράδειγμα**: `"there ,and then"` → `"there, and then"`
- **Παράδειγμα**: `"Hello !"` → `"Hello!"`

#### Κανόνας: Ομάδες Στίξης (... !? κλπ)
- **Περιγραφή**: Ομάδες σημείων στίξης μένουν μαζί, χωρίς κενά εσωτερικά
- **Παράδειγμα**: `"thinking . . ."` → `"thinking..."`
- **Παράδειγμα**: `"Really ! ?"` → `"Really!?"`

### 3.5 Κανόνας Εισαγωγικών (')

#### Μονή Λέξη
- **Περιγραφή**: Τα εισαγωγικά κολλάνε αριστερά και δεξιά της λέξης
- **Παράδειγμα**: `"I am: ' awesome '"` → `"I am: 'awesome'"`

#### Πολλές Λέξεις
- **Περιγραφή**: Τα εισαγωγικά κολλάνε στην πρώτη και τελευταία λέξη
- **Παράδειγμα**: `"' I am the best '"` → `"'I am the best'"`

### 3.6 Κανόνας A/An

#### Κανόνας: a → an
- **Περιγραφή**: Το "a" γίνεται "an" αν η επόμενη λέξη αρχίζει από φωνήεν (a,e,i,o,u) ή h
- **Παράδειγμα**: `"a amazing"` → `"an amazing"`
- **Παράδειγμα**: `"a honest"` → `"an honest"`
- **Παράδειγμα**: `"a book"` → `"a book"` (δεν αλλάζει)

---


# Golden Test Set (Success Test Cases)

## Basic Functional Tests (από Audit Examples)

## Notes on Testing Strategy

1. **Isolation Tests**: Tests 1-6 validate individual rules
2. **Integration Tests**: Tests 7-11 validate rule combinations
3. **Comprehensive Test**: Test 12 validates real-world usage with multiple rules interacting
4. **Edge Cases**: Focus on boundaries (zero values, empty modifiers, consecutive punctuation)
5. **Context Sensitivity**: Tests that validate context awareness (quotes, modifiers with numbers, a/an before modified words)

### Test 1: Mixed Case Transformations
**Input:**
```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6)
```
**Expected Output:**
```
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness
```
**Covers:** (cap), (up), (cap, N) with punctuation

---

### Test 2: Hexadecimal and Binary Conversions
**Input:**
```
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
```
**Expected Output:**
```
Simply add 66 and 2 and you will see the result is 68.
```
**Covers:** (hex), (bin) with punctuation

---

### Test 3: A to An Correction
**Input:**
```
There is no greater agony than bearing a untold story inside you.
```
**Expected Output:**
```
There is no greater agony than bearing an untold story inside you.
```
**Covers:** a/an rule

---

### Test 4: Punctuation Spacing
**Input:**
```
Punctuation tests are ... kinda boring ,what do you think ?
```
**Expected Output:**
```
Punctuation tests are... kinda boring, what do you think?
```
**Covers:** Punctuation spacing, groups of punctuation

---

### Test 5: Quote Handling (Single Word)
**Input:**
```
I am exactly how they describe me: ' awesome '
```
**Expected Output:**
```
I am exactly how they describe me: 'awesome'
```
**Covers:** Single word between quotes

---

### Test 6: Quote Handling (Multiple Words)
**Input:**
```
As Elton John said: ' I am the most well-known homosexual in the world '
```
**Expected Output:**
```
As Elton John said: 'I am the most well-known homosexual in the world'
```
**Covers:** Multiple words between quotes

---

## Advanced Test Cases (Tricky Scenarios)

### Test 7: Multiple Modifiers in Sequence
**Input:**
```
the word (up) was then (cap) followed by another (low) transformation and FF (hex) things.
```
**Expected Output:**
```
the WORD was then Followed by another transformation and 255 things.
```
**Covers:** Multiple different modifiers, hex at end

---

### Test 8: Modifier with Number affecting Multiple Words + Punctuation
**Input:**
```
it was the BEST OF TIMES (low, 3) ! what a story .
```
**Expected Output:**
```
it was the best of times! what a story.
```
**Covers:** (low, N) affecting capitals, multiple punctuation marks

---

### Test 9: Edge Case - A/An with Punctuation and Quotes
**Input:**
```
She found a apple , a orange and a ' honest ' person .
```
**Expected Output:**
```
She found an apple, an orange and an 'honest' person.
```
**Covers:** Multiple a/an corrections, quotes with single word, punctuation

---

### Test 10: Complex Punctuation Groups
**Input:**
```
Wait ... what ! ? Really ? ! I can not believe it ! ! !
```
**Expected Output:**
```
Wait... what!? Really?! I can not believe it!!!
```
**Covers:** Multiple punctuation groups, mixed punctuation

---

### Test 11: Binary/Hex Edge Cases
**Input:**
```
The values 0 (bin) and 0 (hex) are equal , but 1111 (bin) equals F (hex) which is 15 (hex) .
```
**Expected Output:**
```
The values 0 and 0 are equal, but 15 equals 15 which is 15.
```
**Covers:** Zero values, same result from different bases, multiple conversions

---

## Comprehensive Test Paragraph

### Test 12: Large Text with Multiple Rules
**Input:**
```
it (cap) was a amazing DAY (low) ! the sun was shining and the temperature reached 1F (hex) degrees . I went to the store , bought 11 (bin) apples and A (up) orange . the shopkeeper said : ' you are a honest customer ' . when i got HOME (low, 2) , i realized that 101 (bin) plus A (hex) equals F (hex) ! what a DISCOVERY (cap) ... i could not BELIEVE IT (low, 2) ! ? this was the best day EVER (cap, 2) .
```

**Expected Output:**
```
It was an amazing day! the sun was shining and the temperature reached 31 degrees. I went to the store, bought 3 apples and AN orange. the shopkeeper said: 'you are an honest customer'. when i got home, i realized that 5 plus 10 equals 15! what a Discovery... i could not believe it!? this was The Best Day Ever.
```

**Covers:**
- (cap) at start
- a→an with (low) transformed word
- (hex) conversion (1F → 31)
- (bin) conversion (11 → 3)
- (up) on article
- Quote with multiple words
- (low, N) affecting previous words
- Multiple a→an corrections (with h vowel)
- Complex punctuation (... and !?)
- (cap, N) at end of sentence
- Mix of all rules in complex text

---
