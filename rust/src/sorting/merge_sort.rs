#![allow(unused_variables)]
#![allow(dead_code)]
#![allow(unused_mut)]

pub fn merge_sort<T: Ord + Copy>(vals: &mut [T]) {

	if vals.len() < 2 {
		return
	}

	let mid = vals.len() / 2;
	merge_sort(&mut vals[..mid]);
	merge_sort(&mut vals[mid..]);
	merge(vals, mid);
}

fn merge<T: Ord + Copy>(vals: &mut [T], mid: usize) {
	let left_vals = vals[..mid].to_vec();
	let right_vals = vals[mid..].to_vec();

	let mut lidx = 0;
	let mut ridx = 0;

	for v in &mut *vals {
		if ridx == right_vals.len() || (lidx < left_vals.len() && left_vals[lidx] < right_vals[ridx]) {
			*v = left_vals[lidx];
			lidx += 1;
		} else {
			*v = right_vals[ridx];
			ridx += 1;
		}
	}
}

#[cfg(test)]
mod tests {
	use crate::sorting::equal;
	use super::*;

	#[test]
	fn basic() {
		let expected_vals: Vec<isize> = vec![-1, 1, 3, 4, 5];
		let mut vals: Vec<isize> = vec![5, -1, 3, 4, 1];

		merge_sort(&mut vals);
		assert!(equal(&expected_vals, &vals));
	}
}