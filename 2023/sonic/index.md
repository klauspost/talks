<!-- CC BY-SA Klaus Post https://creativecommons.org/licenses/by-sa/4.0/deed.en -->

<!-- 
theme: gaia
class: 
    - lead
-->

<style>
  :root {
    --color-background: #012b35;
    --color-foreground: #edf7f7;
    --color-highlight: #9681c7;
    --color-dimmed: #9681c7;
    --color-dimmed: #95d549;
  }
</style>

# The Feeling of Fast Code

Klaus Post @ Golab 2023

![h:30px](img/minio.png)


---

# 💓 Florence

![bg right:35% brightness:0.75 Chris Yunker](img/florence.jpg)


---

# A story...

![bg right:50%](img/Waterhouse_decameron.jpg)


---

# Our stories are linear

We build our perception from linear events

Cause and Effect

<!-- We arrange them in sequence and learn from cause+effect -->

---

# Parallel Plotlines?

How do you describe a system you are about to build?

<!-- Is it a flowchart? It is a flowchart, right? -->

 ----

# Modern CPUs

A CPU is like a good story. It has several parallel plotlines.

 * Pipelines
 * Branch prediction
 * Out-of-order execution
 * Speculative execution
 * Caching
 * Hyperthreading.


![bg sepia right:30% constrast:200%](img/AMD_Athlon_XP_Thoroughbred_die.jpg)   

---

# 128 concurrent stories?

... all with several parallel plotlines?

<!-- CPU picture... Multiple cores -->
![bg right:35% saturate:50%](img/AMD-Ryzen-5000-Zen-3-Desktop-CPU_Vermeer_Die-Shot_1-1030x679.jpg)

---

# Distributed Systems

* Multiple CPUs
* Multiple nodes
* Remote data
* RPC calls
* Uniformity
* 🤯

![bg right:35% CC-BY Alan Levine](img/network.jpg)

---

# The pieces of the puzzle

* More complexity and data than ever
* Sequential execution doesn't get much better
* The only way to scale is to go parallel
* Design is now more important than ever.

![bg right:35% CC-BY-SA INTVGene](img/puzzle.jpg)


---

# v1 first, optimize later

* Gives a linearly designed system
* Limited window for optimizations
* Optimizations start at the top.

![bg right:33% CC-BY-NC Thomas Hawk](img/v1.jpg)

---

# How do we start thinking about this?

![bg brightness:50% Klaus Post](img/f1-full.jpg)


----
# How to build a race car?

* Performance & Reliability
* Not a solved problem 
* Every decision is a tradeoff
* Money helps, but there is no pay-to-win
* CFD cannot tell you how the car is to drive.

![bg right:35% saturate:75% Klaus Post](img/f1-port2.jpg)

----
# How to build a car?

> **If you get the underlying architecture wrong at the very least you stuck with it for a season.**

*Andrian Newey, F1 podcast*

![bg right:35% saturate:85% AtomsRavelAz](img/Adrian_Newey_2011.jpg)

----

# Building on Faith

Big leaps can require an amount of faith.

* If you think you are more clever than everyone else, you are probably wrong
* ... except like [The Story of the Ferrari 640](https://www.youtube.com/watch?v=QBCwNINnAYY)
* You can't always test big changes before the system is complete
* Convince your most pessimistic colleague
* Be clear about risks and rewards.

![bg right:20% CC-BY-NC Jagrap](img/ildomo.jpg)


---

# Racing Lessons

* You can *always* go faster
* Architecture sets the asymptotic limit
* Testing must correlate with reality.

![bg right:38% saturate:85% Klaus Post](img/podium.jpg)

---
# How to think about speed?

![bg](img/factorio.jpg)

---

# Speeding up...

* Most speedups are specializations
* Speedups seem obvious afterward - creativity
* Focus on proven bottlenecks
* Allocations.

![bg right:38% saturate:85% Klaus Post](img/speed.jpg)

---

# Truly separate work

* Simplify single-thread bound work
* Longest part gives wall speed
* "Lockless" is rarely so
* Sharding breaks single threaded performance
* Experiment with division of work
* Buffers or low latency?

![bg right:38% CC-BY credit: homegets.com](img/pasta.jpg)


---

# Get out of your comfort zone

* Think concurrent
* Consider points that will serialize your code
* Follow the code
* Limit your indexing.

![bg right:33% CC-NC-ND Neal Wellons](img/comfortable.jpg)

---

# Use the tools you have

* Benchmarks `go test -bench=X`
* Profilers  `go tool pprof`
* Tracing `go tool trace`
* Disassembler `go build -gcflags=-S`
* Bounds checks `-gcflags="-d=ssa/check_bce/debug=1"`
* Escape analysis `-gcflags="-m -l"`
* Inlining `-gcflags="-m=2"`

Learn to understand the output of these tools.

![bg right:25% CC-BY-NC Daniel Go](img/tools.jpg)

---

# Myths & Legends

* Branches are slow
* Atomics checks all CPUs
* Fewer instructions are faster
* Memory always go to cache
* Your CPU is incredible (no matter ISA)

![bg right:25% Stable Diffusion Klaus Post](img/unicorns.jpg)


---

# Test (your assumptions)

* Always test your assumptions
* Benchmarks should correlate with real use
* Benchmarks should be reproducible
* Microbenchmarks are often misleading.

![bg right:30% CC-BY-NC U.S. Pacific Fleet](img/hornet.jpg)

---

# Build your dreams

* Keep an open mind
* Feels like a puzzle
* Remember it is a SKILL that builds with PRACTICE.


![bg grayscale:1 right:30% contrast:150%](img/san-matteo.jpg)

---

# Q & A

![Questions](img/questions-tr.png)

---

# Thank You!

Klaus Post @ Golab 2023

`𝕏: @sh0dan`

![h:30px](img/minio.png)

