/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		backgroundColor: '#252525',
		textColor: '#d4d4d4',
		extend: {}
	},

	daisyui: {
		themes: [
			{
				default: {
          error: '#ff0000',
					primary: '#ffde69',
					accent: '#ffee52'
				}
			}
		]
	},

	plugins: [require('daisyui')]
};
