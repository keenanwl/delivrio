export function toNotNullArray<T>(input: Array<T | null | undefined> | undefined): T[] {
	let output: T[] = [];
	input?.reduce((f, val) => {
		const p = val;
		if (!!p) {
			f.push(p);
		}
		return f;
	}, output);
	return output
}
