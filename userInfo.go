package main

import (
	"errors"
	"fmt"
)

type userInfoNode struct {
	userName     string
	password     string
	userBookings []*bookingInfoNode
	left         *userInfoNode
	right        *userInfoNode
}

type BST struct {
	root *userInfoNode
}

var (
	userBst     = &BST{root: nil}
	currentUser *userInfoNode
	userName    string
)

func (bst *BST) validateLogin(userName string, password string) (*userInfoNode, error) {
	if userName == "" || password == "" {
		return nil, errors.New("username or password must not be empty")
	}
	if user, err := userBst.searchUser(userName); err != nil {
		return nil, err
	} else {
		if user.password == password {
			return user, nil
		}
		return nil, errors.New("username or password is invalid")
	}
}

func (bst *BST) insertUserNode(u **userInfoNode, userName string, password string) error {

	if *u == nil {
		newNode := &userInfoNode{
			userName: userName,
			password: password,
			left:     nil,
			right:    nil,
		}
		*u = newNode
		return nil
	}

	if (*u).userName == userName {
		return errors.New("userName is already in use")
	}
	if userName < (*u).userName {
		bst.insertUserNode(&(*u).left, userName, password)
	} else {
		bst.insertUserNode(&(*u).right, userName, password)
	}
	return nil

}

func (bst *BST) createUser(userName string, password string) error {
	bst.insertUserNode(&bst.root, userName, password)
	return nil
}

func (bst *BST) searchForUserNode(u *userInfoNode, userName string) (*userInfoNode, error) {
	if u == nil {
		return nil, errors.New("there are no users")
	} else {
		if u.userName == userName {
			return u, nil
		} else {

			if u.left == nil && u.right == nil {
				return nil, errors.New("user is not found")
			}
			if u.userName > userName {
				return bst.searchForUserNode(u.left, userName)
			} else if u.userName < userName {
				return bst.searchForUserNode(u.right, userName)
			}
		}
	}
	return nil, errors.New("unexpected error has occured")
}

func (bst *BST) searchUser(userName string) (*userInfoNode, error) {
	result, err := bst.searchForUserNode(bst.root, userName)
	return result, err
}

func getNewUserName()string{
	userName := userRawStringInput("\nPlease enter your username: ")
	if user, _ := userBst.searchUser(userName); user != nil {
		fmt.Println(fmt.Errorf("%v is already in use", userName))
		getNewUserName()
	}
	return userName
}

func getNewPassword() string {
	password := userRawStringInput("\nPlease enter your password: ")
	passwordConfirmation := userRawStringInput("\nConfirm password: ")
	if password != passwordConfirmation {
		fmt.Println(errors.New("password confirmation is wrong"))
		getNewPassword()
	}
	return password
}

func (bst *BST) inOrderTraversal(t *userInfoNode) error {
	
	if bst.root == nil {
		panic(errors.New("there are no users"))
	}
	if t != nil {

		bst.inOrderTraversal(t.right)
		fmt.Println("User:", t.userName)
		bst.inOrderTraversal(t.left)
	}
	return nil
}

func (bst *BST) printAllUsers() {
	bst.inOrderTraversal(bst.root)
}

func (bst *BST) findSuccessor(t *userInfoNode) *userInfoNode {
	for t.right != nil {
		t = t.right
	}
	return t
}

func (bst *BST) removeNode(t **userInfoNode, userName string) (*userInfoNode, error) {
	
	if *t == nil {
		return nil, errors.New("there are no users")
	}else if userName < (*t).userName {
		(*t).left, _ = bst.removeNode(&(*t).left, userName)
	}else if userName > (*t).userName {
		(*t).right, _ = bst.removeNode(&(*t).right, userName)
	}else { 
		if (*t).left == nil && (*t).right == nil{ //no child
			*t = nil
			return *t, nil
		}else if (*t).left == nil{
			temp := *t
			*t = temp.right
			temp = nil
			return *t, nil
		}else if (*t).right == nil{
			temp := *t
			*t = temp.left
			temp = nil
			return *t, nil
		}else {
			temp:= bst.findSuccessor((*t).right)
			(*t) = temp
			(*t).right, _ = bst.removeNode(&(*t).right, userName)
		}
	}
	return *t, nil
}

func (bst *BST) deleteUser(userName string) {
	bst.root, _ = bst.removeNode(&bst.root, userName)
}