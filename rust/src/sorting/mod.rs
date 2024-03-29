mod insertion_sort;
mod bubble_sort;
mod merge_sort;

pub use self::insertion_sort::insertion_sort;
pub use self::bubble_sort::bubble_sort;
pub use self::merge_sort::merge_sort;

use std::cmp;

pub fn equal<T: cmp::PartialOrd + std::fmt::Debug>(expected_vals: &[T], vals: &[T]) -> bool {
	if expected_vals.len() != vals.len() {
		return false;
	}

	for i in 0..expected_vals.len() {
		let expected = expected_vals.get(i);
		let val = vals.get(i);
		if expected != val {
			println!("Expected {:#?} but got {:#?} instead", expected, val);

			return false;
		}
	}

	return true;
}