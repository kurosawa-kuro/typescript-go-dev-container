type Micropost = {
    id: number;
    title: string;
    created_at: string;
    updated_at: string;
  };
  
  type Props = {
    initialMicroposts: Micropost[];
  };
  
  export default function MicropostList({ initialMicroposts }: Props) {
    return (
      <div className="space-y-4">
        {initialMicroposts.map((post) => (
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
    );
  }