import colors from 'tailwindcss/colors';
import starlightPlugin from '@astrojs/starlight-tailwind';

/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}'],
  theme: {
    extend: {
      colors: {
        // Accent color. Indigo is closest to Starlight’s defaults.
        accent: colors.indigo,
        // Gray scale. Zinc is closest to Starlight’s defaults.
        gray: colors.zinc,
      },
    },
    fontFamily: {
    	// Text font.
    	sans: ['"Source Sans Pro"'],
      // Code font.
      mono: ['"IBM Plex Mono"'],
    },
  },
  plugins: [starlightPlugin()],
};