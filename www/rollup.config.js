import svelte from 'rollup-plugin-svelte';
import resolve from 'rollup-plugin-node-resolve';
import commonjs from 'rollup-plugin-commonjs';
import buble from 'rollup-plugin-buble';
import { uglify } from 'rollup-plugin-uglify';

const production = !process.env.ROLLUP_WATCH;

export default {
	input: 'src/index.js',
	output: {
		file: 'public/bundle.js',
		format: 'iife',
		name: 'hurriyetlite',
		sourcemap: false,
	},
	plugins: [
		svelte({
			css: css => {
				css.write('public/bundle.css', false);
			},
			cascade: false
		}),

		resolve(),
		commonjs(),

		production && buble({ exclude: 'node_modules/**' }),
	]
};
