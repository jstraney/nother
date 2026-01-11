"use client";

interface TreeListProps {
  children: React.ReactNode;
}

export function TreeList({ children }: TreeListProps) {
  return <ul className="space-y-1">{children}</ul>;
}

interface TreeListItemProps {
  children: React.ReactNode;
}

export function TreeListItem({ children }: TreeListItemProps) {
  return <li className="[&>ul]:pl-4">{children}</li>;
}
