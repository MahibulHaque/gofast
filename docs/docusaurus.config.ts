import {themes as prismThemes} from 'prism-react-renderer';
import type {Config} from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';

const config: Config = {
  title: 'Gofast',
  tagline:
    'Effortlessly generate the ideal application scaffold for your Go web application or API, customized to your needs and saving yourself hours of time and thinking.',
  favicon: 'img/favicon.ico',

  future: {
    v4: true,
  },

  url: 'https://mahibulhaque.github.io',
  baseUrl: '/gofast/',
  trailingSlash: false,

  organizationName: 'mahibulhaque',
  projectName: 'gofast',
  deploymentBranch: 'gh-pages',

  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',

  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      {
        docs: {
          sidebarPath: './sidebars.ts',
        },
        theme: {
          customCss: './src/css/custom.css',
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    image: 'img/docusaurus-social-card.jpg',
    navbar: {
      title: 'Gofast',
      logo: {
        alt: 'My Site Logo',
        src: 'img/logo.svg',
      },
      items: [
        {
          type: 'docSidebar',
          sidebarId: 'tutorialSidebar',
          position: 'left',
          label: 'Docs',
        },
        {
          href: 'https://github.com/mahibulhaque/gofast',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Learn',
          items: [
            {label: 'Introduction', to: '/docs/intro'},
            {label: 'Installation', to: '/docs/installation'},
          ],
        },
        {
          title: 'Connect with me',
          items: [
            {
              label: 'LinkedIn',
              href: 'https://www.linkedin.com/in/mahibulhaque/',
            },
            {
              label: 'X',
              href: 'https://x.com/Mahibul45291325',
            },
          ],
        },
        {
          title: 'More',
          items: [
            {label: 'GitHub', href: 'https://github.com/mahibulhaque/gofast'},
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()}. Built with ❤  by Mahibul Haque`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
