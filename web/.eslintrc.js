module.exports = {
    'env': [
        'plugin:vue/strongly-recommended',
        'standard',
    ],

    'extends': [
        'plugin:vue/essential',
        'standard',
    ],

    'globals': {
        'Atomics': 'readonly',
        'SharedArrayBuffer': 'readonly',
    },

    'parserOptions': {
        'ecmaVersion': 2018,
        'sourceType': 'module',
    },

    'plugins': [
        'vue',
    ],

    'rules': {
        'indent': ['error', 4],
    },

    env: {
        browser: true,
        es6: true,
    },

    globals: {
        Atomics: 'readonly',
        SharedArrayBuffer: 'readonly',
    },

    parserOptions: {
        ecmaVersion: 2018,
        sourceType: 'module',
    },

    plugins: [
        'vue',
        'typescript',
    ],

    rules: undefined,
};
