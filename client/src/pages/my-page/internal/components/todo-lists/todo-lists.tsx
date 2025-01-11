import { serverFetch } from "../../../../../utils/fetch";

interface TodoListsProps {
    userID: number;
}

export const TodoLists = async ({ userID }: TodoListsProps) => {
    const todoLists = await serverFetch(`/api//todolists?user_id=${userID}`);
    console.log(todoLists);
    return(
        <div>
            ここにリスト
        </div>
    )
}