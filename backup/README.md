# Backup

The backup module allows orders to be saved to or loaded from file. Orders are saved as boolean values where active orders are ``true`` and no order is ``false``. The priority queue is the current prioritized floors to go to.

### Saved data:

- Cab calls
- Hall calls
- Priority queue

### Backup file structure:

cab: []bool

hall: [][]bool

priority queue: []int