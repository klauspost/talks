<!-- CC BY-SA Klaus Post https://creativecommons.org/licenses/by-sa/4.0/deed.en -->

<!-- 
theme: gaia
class: 
    - lead
paginate: false
-->

<!-- size: 16:9 -->

<style>
  :root {
    --color-background: #012b35;
    --color-foreground: #edf7f7;
    --color-highlight: #9681c7;
    --color-dimmed: #95d549;
  }
</style>

# Go Data Compression

Klaus Post @ Golab 2026

![h:30px](img/minio.png)

Slides: https://golab2026.klauspost.com

`100% meat-generated (except stats)`

---

![bg h:90%](img/memes/02.png)

---

# Data Compression Primer

Minimize transfer and storage

Always a CPU/size tradeoff

Sometimes Compatibility 

---

# Generic Compression
# Domain Specific (+lossy)

---

![bg h:90%](img/memes/08.png)

---

# Match Coding

String of Operations

Literals and back-references

Main compression scheme

[Example](block-vis/index.html)

---

# Entropy Coding

Symbols (bytes) -> Bit Stream

* Huffman
* Range Coding
* ANS

Not order dependent

---

# DEFLATE

Create by Phil Katz for PKZIP

![sepia Phil Katz - https://en.wikipedia.org/wiki/File:Phil_Katz.jpg](img/Phil_Katz.jpg)

Huffman Stream with either a literal or a match.

---

![bg](img/deflate-buffers.png)


---

# Go 1.27 Improvements

- Custom level 1-6 encoders
- Huffman Encoding
- Level 7-9 tweaks

[NUMBERS pls](http://127.0.0.1:8080/#c=silesia&m=1omrle93&t=1&by=lib&lb=std%40go1.26.10,std%40go1.27.1)

---

# Stdlib Zstandard

WIP - slow but steady.

---

![bg](img/zstd-buffers.png)

---

# Transformers 

(no, not those)

Usually doesn't compress by itself.

Example: Columnar data

---

# Context Mixing / LLM

Select Encoder based on history.

Good. Exotic. (Too) Slow.

---

# Benchmarking

![h:50%](img/memes/03.jpg)

---

# Benchmarking

Dimensions:
* Speed
* Compression Ratio
* Memory Usage
* Comp vs Decomp
* Representative test data
* CPU/memory

---

# [Let's see some](http://127.0.0.1:8080/#c=silesia&m=1omrle93&t=1&op=c)

---

![w:800px Threads Compression](./img/threads-comp.png)

[Threads Compression](http://127.0.0.1:8080/#c=apache-log&m=1omrle93&t=1,2,4,8,12,16&by=threads&cd=minlz)

---

![w:800px Threads Decompression](./img/threads-decomp.png)

[Threads Decompression](http://127.0.0.1:8080/#c=apache-log&m=1omrle93&t=1,2,4,8,12,16&by=threads&cd=minlz)

---

![w:800px Window Size Compression](./img/window-comp.png)

Window Size Compression

---

![w:800px Window Size Decompression](./img/window-decomp.png)

Window Size Decompression


# Third Party - Byte Aligned

| Compressor | Main Upsides                | Main Downside |
|------------|-----------------------------|---------------|
| Snappy     | +Compatible, OK Speed       | Low Comp      |
| LZ4        | +Compatible, +Speed, Simple | Low Comp      |
| MinLZ      | +Speed, Backcompat          | Medium Comp   |

---

# Third Party - Bit Aligned

| Compressor     | Main Upsides                      | Main Downside               |
|----------------|-----------------------------------|-----------------------------|
| gzip (deflate) | +Compatible                       | Slow Decomp, Medium Comp    |
| bzip2          | Compatible                        | Very Slow, Including Decomp |
| zstd           | Compatible, Size, Speed, Versaile | Good Decomp Speed           |
| brotli         | HTML, Size                        | Rarely the best             |


---

# MinLZ SIMD Blocks

Design a compressor purely around SIMD operations

Goal: Entropy coding with minimal extra CPU

Achieved: Entropy coding with 2x-3x faster decompression

---

![bg](img/simd-block-buffers.png)

---


# Q & A

![bg h:80% right:35% saturate:75%](img/memes/09.png)

![Questions](img/questions-tr.png)

---

# Thank You!

Klaus Post @ Golab 2026

`𝕏: @sh0dan`
`🦋: @klauspost.com`

![h:30px](img/minio.png)

