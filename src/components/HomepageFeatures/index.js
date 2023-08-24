import React from 'react';
import clsx from 'clsx';
import styles from './styles.module.css';

const FeatureList = [
  {
    title: 'Rancher Extension',
    Svg: require('@site/static/img/rancher-logo-cow-blue.svg').default,
    description: (
      <>
        Designed to extend and fully intergrate with Rancher.
      </>
    ),
  },
  {
    title: 'Full Cluster API Support',
    Svg: require('@site/static/img/capi_logo.svg').default,
    description: (
      <>
        Use any provider you want and still have your cluster available within Rancher Manager.
      </>
    ),
  },
];

function Feature({ Svg, title, description }) {
  return (
    <div className={clsx('col col--6')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
