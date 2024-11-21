type Micropost = {
  id: number;
  title: string;
  created_at: string;
  updated_at: string;
};

async function getMicroposts(): Promise<Micropost[]> {
  const res = await fetch('http://backend:8000/microposts', {
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
      <h1 className="text-2xl font-bold text-gray-900 dark:text-white mb-6">
        Microposts
      </h1>
      <div className="space-y-4">
        {microposts.map((post) => (
          <div 
            key={post.id}
            className="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-sm"
          >
            <h2 className="text-xl font-semibold text-gray-900 dark:text-white">
              {post.title}
            </h2>
            <div className="mt-2 text-sm text-gray-500 dark:text-gray-400">
              Created: {new Date(post.created_at).toLocaleDateString()}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}