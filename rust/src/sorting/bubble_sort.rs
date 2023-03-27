pub fn bubble_sort<T: Ord + Copy>(vals: &mut [T]) {
	if vals.is_empty() {
		return;
	}

	loop {
		let mut swapped = false;
		for i in 1..vals.len() {
			let cur = vals[i];

			if vals[i-1] > cur {
				vals.swap(i, i-1);

				swapped = true;
			}
		}

		if !swapped {
			break;
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

		bubble_sort(&mut vals);
		assert!(equal(&expected_vals, &vals));
	}
}