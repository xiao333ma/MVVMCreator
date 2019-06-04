package mvvm

import "MVVMCreator/mvvm/utils"

type ViewModelCreator struct {
	Creator
}

func (r ViewModelCreator) CreateDotHFile()  {
	str := `//
//  __prefix__ViewModel.h
//  __proj__
//
//  Created by __user__ on __time__.
//  Copyright © __year__ __user__. All rights reserved.
//

#import "IGCViewModel.h"

NS_ASSUME_NONNULL_BEGIN

@interface __prefix__ViewModel : IGCViewModel

@end

NS_ASSUME_NONNULL_END
`
	utils.WriteFile(utils.FileTypeOfViewModel, true, str)

}

func (r ViewModelCreator) CreateDotMFile()  {

	str :=`//
//  __prefix__ViewModel.m
//  __proj__
//
//  Created by __user__ on __time__.
//  Copyright © __year__ __user__. All rights reserved.
//

#import "__prefix__ViewModel.h"
#import "__prefix__Pipeline.h"

@interface __prefix__ViewModel ()

@property (nonatomic, strong) __prefix__Pipeline *pipeline;

@end

@implementation __prefix__ViewModel

- (instancetype)initWithContext:(NSDictionary *)context {
    self = [super initWithContext:context];
    if (self) {
        
    }
    return self;
}

- (void)setupSubViewModelByContext:(NSDictionary *)context {
    
}

- (void)addObservers {
    
}

- (void)fetchData {
    
}

#pragma mark - Accessor

- (__prefix__Pipeline *)pipeline {
    if (!_pipeline) {
        _pipeline = [[__prefix__Pipeline alloc] init];
    }
    return _pipeline;
}

@end
`
	utils.WriteFile(utils.FileTypeOfViewModel, false, str)
}