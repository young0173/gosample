#! /usr/bin/python
# -*- coding: utf-8 -*-

import os
import sys
import argparse
import shutil


def rm_recursive(dir_rm, file_rm, base_dir):

	if not os.path.isdir(base_dir):
		return False

	try:
		for r in os.listdir(base_dir):

			scan_dir = base_dir + "/" + r
			print(scan_dir)

			if not os.path.isdir(scan_dir):
				if r == file_rm:
					os.remove(scan_dir)
					print("removing file:" + scan_dir)
				continue
			elif r == dir_rm:
				shutil.rmtree(scan_dir)
				print("removing dir	:" + scan_dir)
			else:
				rm_recursive(dir_rm, file_rm, scan_dir)
		return True
	except Exception as e:
		print(e)
		return False






if __name__ == '__main__':

	base_dir = sys.path[0]

	argparser = argparse.ArgumentParser(description='input argument')
	argparser.add_argument('--dirname', type=str, default=None)
	argparser.add_argument('--filename', type=str, default=None)
	args = argparser.parse_args()

	dir_rm = args.dirname
	file_rm = args.filename

	ok = rm_recursive(dir_rm, file_rm, base_dir)
	if ok:
		print("Delete dir successfully.")
	else:
		print("Delete dir failed.")
