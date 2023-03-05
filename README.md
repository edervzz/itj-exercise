# itj-code-exercise
Platform Science Code Exercise

## Build and Run

1. Clone this repository on your personal computer, open a terminal and move to your project local folder
2. To build, run in your terminal `go build`
3. To run execute file `itj-code-exercise`
4. Follow next instructions from executable.

## Aim

Assigns shipment destinations to drivers in a way that maximizes the total suitability score (SS) over the set
of drivers.

## Constraints
Each driver can only have one shipment and each shipment can only be offered to one driver.

## Business rules:
1. If the length of the shipment's destination street name is even, 
the base SS is the number of vowels in the driver’s name multiplied by 1.5.
2. If the length of the shipment's destination street name is odd, 
the base SS is the number of consonants in the driver’s name multiplied by 1.
3. If the length of the shipment's destination street name shares any common factors (besides 1) with the length of the driver’s name, 
the SS is increased by 50% above the base SS.

## Assumptions

- Address particles are: `number, street, inner number, city, state, zipcode`.
- Number and street are just separated by space.
- `["Suite", "Apt."]` are not part of street name but the inner number.
- e.g.: 63187 Volkman Garden Suite 447, San Diego, CA 92126

## Example
### Driver's name: Daniel Davidson
where it has a length = 15, 6 vowels, 8 consonants.
### Shipment's destination: 44 Fake Dr., San Diego, CA 92122
where street name is "Fake Dr." with a lenght = 8, (even)
### Then
1. choose the rule 1. "the length of the shipment's destination street name is even" due to street name length is 8
 
2. use the number of vowel in the driver’s name (6) and multiply by 1.5
### Result
9

## Hungarian Algorithm
This Algorithm create the assignment between two set of objects where we want the best assignment depending efficiency. 
In our case efficiency depends of suitability score (SS).

This response to O(n^3) complexity because on every single loop-calculation we going to get a new matrix.
