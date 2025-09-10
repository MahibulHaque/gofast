import type {ReactNode} from 'react';
import Heading from '@theme/Heading';
import styles from './styles.module.css';

type FeatureItem = {
  title: string;
  description: ReactNode;
};

const FeatureList: FeatureItem[] = [
  {
    title: 'Cut out time-consuming setup work',
    description: (
      <>
Gofast creates a fully-functioning application scaffold for you, with all the setup steps and package integrations taken care of. This means you can <strong>
get started fast and focus on your application-specific code
        </strong>, instead of the boring boilerplate.
      </>
    ),
  },
  {
    title: 'Tested Application Structure',
    description: (
      <>
The generated code gives you a really solid foundation to continue building on, with a structured application, minimal complexity and idomatic golang code.
      </>
    ),
  },
  {
    title: 'You fully control the code',
    description: (
      <>
        Gofast is not a third-party framework, rather a minimal CLI which only serves the purpose of scaffolding a new golang repository for you. The extension of the starter application fully relies on your preference and choice.
      </>
    ),
  },
  {
    title: 'Only the required features',
    description: (
      <>
You get to customize your application scaffold to include only the features that you need, which means there will be fewer dependencies resulting in a smaller binary.
      </>
    ),
  },
];

function Feature({title, description}: FeatureItem) {
  return (
    <div className={styles.feature}>
        <h3 className={styles.featureHeader}>{title}</h3>
        <p>{description}</p>
    </div>
  );
}

export default function HomepageFeatures(): ReactNode {
  return (
    <section>
      <div className={styles.outerContainer}>
        <div className={styles.featureContainer}>
          {FeatureList.map((feature)=>(
          <Feature key={feature.title} title={feature.title} description={feature.description}/>
          ))}
        </div>
      </div>
    </section>
  );
}
