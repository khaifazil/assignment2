package main

import (
	"errors"
)

type userInfoNode struct {
	userName string
	password string
	userBookings []*bookingInfoNode
	left     *userInfoNode
	right    *userInfoNode
}

type BST struct {
	root *userInfoNode
}

var userBst = &BST{root: nil}

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


