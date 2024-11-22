'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import MicropostModal from '@/app/components/MicropostModal';

export default function CreateMicropostButton() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  const handleCreateMicropost = async (title: string) => {
    setIsLoading(true);
    setError(null);

    try {
      const response = await fetch('http://localhost:8000/api/microposts', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ title }),
        credentials: 'include',
      });

      // レスポンスのステータスとContent-Typeをチェック
      const contentType = response.headers.get('content-type');
      if (!response.ok) {
        if (contentType && contentType.includes('application/json')) {
          const errorData = await response.json();
          throw new Error(errorData.error || 'Failed to create micropost');
        } else {
          const textError = await response.text();
          throw new Error(textError || `HTTP error! status: ${response.status}`);
        }
      }

      // 成功レスポンスの処理
      if (contentType && contentType.includes('application/json')) {
        const data = await response.json();
        console.log('Created micropost:', data);
      }

      // 投稿成功後の処理
      router.refresh();
      setIsModalOpen(false);
    } catch (error) {
      console.error('Error creating micropost:', error);
      setError(error instanceof Error ? error.message : 'An unexpected error occurred');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <>
      <button
        onClick={() => setIsModalOpen(true)}
        className="px-4 py-2 text-sm font-medium text-white bg-blue-600 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50"
        disabled={isLoading}
      >
        {isLoading ? 'Creating...' : 'New Post'}
      </button>

      <MicropostModal
        isOpen={isModalOpen}
        onClose={() => !isLoading && setIsModalOpen(false)}
        onSubmit={handleCreateMicropost}
      />
    </>
  );
}