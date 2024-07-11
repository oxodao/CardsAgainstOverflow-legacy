module.exports = {
    root: true,
    env: {
        node: true,
    },
    extends: [
        'plugin:vue/recommended', // flat/vue2-recommended
        'eslint:recommended',
    ],
    parserOptions: {
        ecmaVersion: 2020,
    },
    rules: {
        'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
        'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
        'indent': ['error', 4],
        'vue/multi-word-component-names': 'off',
        'vue/script-indent': ['error', 4, {
            'baseIndent': 0,
        }],
        'vue/html-self-closing': ['error', {
            'html': {
                'void': 'always',
                'normal': 'always',
                'component': 'always'
            },
            'svg': 'always',
            'math': 'always'
        }],
        'vue/html-indent': ['error', 4, {
            'attribute': 1,
            'baseIndent': 1,
        }],
        'quotes': [2, 'single', {'avoidEscape': true}],
        'semi': [2, 'always'],
        'keyword-spacing': ["warn", {"before": true}],
        'curly': 'error'
    },
    overrides: [
        {
            files: [
                '**/__tests__/*.{j,t}s?(x)',
                '**/tests/unit/**/*.spec.{j,t}s?(x)',
            ],
            env: {
                mocha: true,
            },
        },
        {
            'files': ['*.vue'],
            'rules': {
                'indent': 'off'
            }
        }
    ],
}