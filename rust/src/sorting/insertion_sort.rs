pub fn insertion_sort<T: Ord + Copy>(vals: &mut [T]) {
	
	for i in 1..vals.len() {
		let key = vals[i];
		let mut j = i;

		while j > 0 && vals[j - 1] > key {
			vals[j] = vals[j - 1];
			j -= 1;
		}

		vals[j] = key;
	}
}

#[cfg(test)]
mod tests {
	// importing names from outer (for mod tests) scope
	use crate::sorting::equal;
	use super::*;

	#[test]
	fn basic() {
		let expected_vals: Vec<isize> = vec![-1, 1, 3, 4, 5];
		let mut values: Vec<isize> = vec![5, -1, 3, 4, 1];
		insertion_sort(&mut values);
		println!("{:?}", values);

		assert!(equal(&expected_vals, &values));
	}
}