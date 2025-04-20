# Fisher-Yates Shuffle (Durstenfeld Variation)

The Fisher-Yates algorithm efficiently shuffles a sequence, guaranteeing a uniform distribution of permutations. This is the standard method for unbiased shuffling.

## Algorithm

Iterate from the first element up to the second-to-last element (index `i`):

1. Choose a random index `j` such that `i <= j < n` (where `n` is the sequence length).
2. Swap the elements at indices `i` and `j`.

**Why this works (Unbiasedness):**

The key is that at each step `i`, every element *not yet placed* in its final position (i.e., those at indices `i` through `n-1`) has an equal probability, `1 / (n - i)`, of being selected by the random choice `j` and swapped into position `i`. This ensures that every element has an equal chance of ending up in any position.

Consider shuffling `[A, B, C]` (`n=3`):

- **Step `i = 0`**: `j` is chosen uniformly from `[0, 1, 2]`. Range size `n-i = 3`. Any of `A`, `B`, `C` has a `1/3` chance of being swapped to index 0.
  - Example: `j=1` is chosen (prob `1/3`). Swap `A` and `B`. Array becomes `[B, A, C]`.
- **Step `i = 1`**: `j` is chosen uniformly from `[1, 2]`. Range size `n-i = 2`. The elements currently at indices 1 and 2 (`A`, `C`) each have a `1/2` chance of being swapped to index 1.
  - Example: `j=2` is chosen (prob `1/2`). Swap `A` and `C`. Array becomes `[B, C, A]`.
- **Step `i = 2` (n-1)**: Loop finishes.

The final permutation `[B, C, A]` occurred with probability `(1/3) * (1/2) = 1/6`. A similar calculation shows that *every* possible permutation (`n! = 3! = 6` of them) has exactly a `1/n!` probability of occurring, which is the definition of a uniform shuffle.

(Implementation: [`deck.go`](../../legacy/cards/deck.go#L77-L83))

```go
import "math/rand"

// shuffleWithFY shuffles the deck in place using the Fisher-Yates algorithm.
// Ensures a uniform random permutation.
func (d deck) shuffleWithFY() {
    n := len(d)
    // Iterate backwards from the last element down to the second element.
    // Or iterate forwards from the first element up to the second-to-last.
    // Both are valid implementations of Fisher-Yates. This example iterates forwards.
    for i := 0; i < n-1; i++ {
        // Choose random index j from the remaining unshuffled portion (i to n-1)
        j := rand.Intn(n-i) + i // j is in the range [i, n-1]
        // Swap element i with element j
        d[i], d[j] = d[j], d[i]
    }
}
```

## Common Biased Shuffle (Incorrect)

A frequent mistake is to choose the random index `j` from the *entire* range `[0, n-1]` (or `[0, n-2]`) in each iteration.

```go
// shuffleNaive attempts to shuffle but introduces bias. (DO NOT USE)
func (d deck) shuffleNaive() {
    n := len(d)
    for i := range d {
        // Incorrect: samples j from [0, n-1] instead of [i, n-1]
        // Or even worse, [0, n-2] as shown in the original example.
        j := rand.Intn(n) // Or rand.Intn(n-1) - both are biased!
        d[i], d[j] = d[j], d[i]
    }
}
```

This naive approach **does not** produce a uniform distribution. The probability of certain permutations becomes skewed. For example, with `j := rand.Intn(n-1)`, the last element `d[n-1]` is never directly selected for the swap target `j`.

## Recommendation: Use `rand.Shuffle`

Go's standard library provides `rand.Shuffle`, which implements the Fisher-Yates algorithm correctly and efficiently. **Prefer using `rand.Shuffle`**.

```go
import "math/rand"

// shuffleWithStdLib shuffles the deck using the standard library's `rand.Shuffle`.
// This is the recommended approach.
func (d deck) shuffleWithStdLib() {
    rand.Shuffle(len(d), func(i, j int) {
        d[i], d[j] = d[j], d[i]
    })
}
```

(Implementation: [`deck.go`](../../legacy/cards/deck.go#L88-L92))
