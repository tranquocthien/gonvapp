import Image from 'next/image';
import styles from './img-Overlay.module.css';

export default function ImgOverlay({
  srcImg,
  title,
  description,
  view,
}) {
  return (
    <div className={styles.container}>
      <div
        style={{
          borderRadius: '0.75rem',
          overflow: 'hidden',
          height: '300px',
          aspectRatio: '3/4',
        }}
      >
        <Image
          src={srcImg}
          alt=""
          layout="fill"
          objectFit="cover"
        />
      </div>

      <div className={styles.textBlock}>{view}</div>
      <div className={styles.overlay}>
        <div className={styles.text}>
          <h5>{title}</h5>
          <p>{description}</p>
        </div>
      </div>
      <div className={styles.titleBlock}>
        <h5>{title}</h5>
      </div>
      <p className={styles.onSmallDevices}>{description}</p>
    </div>
  );
}
