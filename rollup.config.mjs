
// rollup.config.js
import resolve from '@rollup/plugin-node-resolve';

export default {
  input: 'index.js',
  output: {
    file: 'static/js/bundle.js',
    format: 'iife'
  },
  plugins: [resolve()]
};
