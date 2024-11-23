import MicropostList from './MicropostList';
import CreateMicropostButton from './CreateMicropostButton';

type Micropost = {
  id: number;
  title: string;
  image_path: string;
  created_at: string;
  updated_at: string;
};

async function getMicroposts(): Promise<Micropost[]> {
  const res = await fetch('http://localhost:8000/api/microposts', {
    headers: {
      'Content-Type': 'application/json',
    },
    cache: 'no-store',
  });

  if (!res.ok) {
    throw new Error('Failed to fetch microposts');
  }

  return res.json();
}

export default async function Home() {
  const microposts = await getMicroposts();

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-2xl font-bold text-gray-900 dark:text-white">
          Microposts
        </h1>
        <CreateMicropostButton />
      </div>
      <MicropostList initialMicroposts={microposts} />
    </div>
  );
}