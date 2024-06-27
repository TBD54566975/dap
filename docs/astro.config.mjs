import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';
import tailwind from '@astrojs/tailwind';

// https://astro.build/config
export default defineConfig({
  site: 'https://daps.dev',
  integrations: [
    starlight({
      title: 'DAP',
      logo: {
        light: '/src/assets/logo-light.svg',
        dark: '/src/assets/logo-dark.svg',
        replacesTitle: true,
      },
      social: {
        github: 'https://github.com/TBD54566975/dap',
      },
      defaultLocale: 'root',
      locales: {
        root: { label: 'English', lang: 'en' },
        es: { label: 'Español', lang: 'es' },
      },
      sidebar: [
        {
          label: 'Introduction',
          link: '/',
          translations: {
            es: 'Introducción',
          },
        },
        {
          label: 'Getting Started',
          link: '/getting-started/',
          translations: {
            es: 'Primeros pasos',
          },
        },
        {
          label: 'Tutorials',
          translations: {
            es: 'Tutoriales',
          },
          items: [
            {
              label: 'Paying a DAP',
              link: '/tutorials/dap-payment/',
              translations: {
                es: 'Pagar un DAP',
              },
            },
            {
              label: 'Adding DAPs to Your App',
              link: '/tutorials/app-integration/',
              translations: {
                es: 'Agregar DAPs a tu aplicación',
              },
            },
          ],
        },
        {
          label: 'Core Concepts',
          translations: {
            es: 'Conceptos Fundamentales',
          },
          items: [
            {
              label: 'DAP Building Blocks',
              link: '/concepts/building-blocks/',
              translations: {
                es: 'Componentes básicos de DAP',
              },
            },
            {
              label: 'Why DAPs?',
              link: '/concepts/why-daps/',
              translations: {
                es: '¿Por qué DAPs?',
              },
            },
            {
              label: 'Comparing Alternatives',
              link: '/concepts/comparing-alternatives/',
              translations: {
                es: 'Comparando alternativas',
              },
            },
            {
              label: 'Empowering Users',
              link: '/concepts/empowering-users/',
              translations: {
                es: 'Empoderando a los usuarios',
              },
            },
            {
              label: 'Privacy Considerations',
              link: '/concepts/privacy/',
              translations: {
                es: 'Consideraciones de privacidad',
              },
            },
          ],
        },
        {
          label: 'How do I...?',
          translations: {
            es: '¿Cómo puedo...?',
          },
          items: [
            {
              label: 'Run a DAP Registry',
              link: '/how-to/dap-registry/',
              translations: {
                es: 'Desplegar un registro de DAP',
              },
            },
            {
              label: 'Register a Custom DID',
              link: '/how-to/custom-did/',
              translations: {
                es: 'Registrar un DID personalizado',
              },
            },
            {
              label: 'Migrate Users to DAPs',
              link: '/how-to/migrate-to-daps/',
              translations: {
                es: 'Migrar usuarios a DAPs',
              },
            },
            {
              label: 'Create a Money Address Format',
              link: '/how-to/create-money-address-format/',
              translations: {
                es: 'Crear un formato de dirección de dinero',
              },
            },
          ],
        },
        {
          label: 'Reference',
          translations: {
            es: 'Referencia',
          },
          items: [
            {
              label: 'DAP Specification',
              link: '/reference/dap-spec/',
              translations: {
                es: 'Especificación de DAP',
              },
            },
            {
              label: 'Money Address Format',
              link: '/reference/money-address/',
              translations: {
                es: 'Formato de dirección de dinero',
              },
            },
            {
              label: 'DAP Registry API',
              link: '/reference/registry-api/',
              translations: {
                es: 'API de registro de DAP',
              },
            },
            {
              label: 'DID Service Extensions',
              link: '/reference/did-services/',
              translations: {
                es: 'Extensiones de servicio DID',
              },
            },
          ],
        },
      ],
      customCss: [
        // Path to daps.dev Tailwind base styles:
        './src/tailwind.css',
        '@fontsource/source-sans-pro',
        '@fontsource/ibm-plex-mono',
      ],
    }),
    tailwind({
      // Disable the default base styles:
      applyBaseStyles: false
    }),
  ],
});