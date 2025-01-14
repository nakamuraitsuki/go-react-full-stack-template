import { serverFetch } from "../../../../../utils/fetch";

interface TodoListsProps {
    userID: number;
}

export const TodoLists = ({ userID }: TodoListsProps) => {
    const todoLists = serverFetch(`/api/todolists?user_id=${userID}`);
    
    return(
        <div>
            ここにリスト
        </div>
    )
}