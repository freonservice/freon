const path = require('path');
const root = path.join(__dirname);

const config = {
    rootDir: root,
    // Targets ========================================================
    serveDir: path.join(root, '.serve'),
    distDir: path.join(root, 'dist'),
    prodDir: path.join(root, 'prod'),
    clientManifestFile: 'manifest.webpack.json',
    clientStatsFile: 'stats.webpack.json',

    // Source Directory ===============================================
    srcDir: path.join(root, 'app'),
    srcServerDir: path.join(root, 'server'),

    // HTML Layout ====================================================
    srcHtmlLayout: path.join(root, 'app', 'index.html'),

    // Site Config ====================================================
    siteTitle: 'Freon',
    siteDescription: 'Freon Admin Panel',
    siteCannonicalUrl: 'http://localhost:4100',
    scssIncludes: [],

    apiServerURL: 'http://localhost:4000',

    module: {
        rules: [
            {
                test: /\.css$/i,
                use: ['style-loader', 'css-loader'],
            },
        ],
    },
};

module.exports = config;