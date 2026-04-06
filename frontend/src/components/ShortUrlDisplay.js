export default function ShortUrlDisplay({ code }) {
    return (
      <div>
        <p>Your short URL:</p>
        <a href={`/${code}`}>{process.env.NEXT_PUBLIC_BASE_URL}/{code}</a>
      </div>
    );
  }